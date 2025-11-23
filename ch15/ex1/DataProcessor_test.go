package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type RemoteCalc struct {
	CalcURL string
	Client *http.Client
}

func (rc RemoteCalc) Calculate(expression string) (string, error) {
	r := strings.NewReader(expression)
	req, err := http.NewRequest(http.MethodPost, rc.CalcURL, r)
	if err != nil {
		return "", err
	}

	res, err := rc.Client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Test_parser(t *testing.T) {
	data := []struct {
		name string
		input []byte
		expected Input
	}{
		{"valid", []byte("CALC_1\n+\n3\n2"), Input{Id: "CALC_1", Op: "+", Val1: 3, Val2: 2}},
		{"invalid", []byte("CALC_1\n+\n3\na"), Input{Id: "", Op: "", Val1: 0, Val2: 0}},
		{"invalid", []byte("CALC_1\n+\na\n2"), Input{Id: "", Op: "", Val1: 0, Val2: 0}},
	}

	for _, d := range data {
		t.Run(d.name, func (t *testing.T) {
			parsedData, _ := parser(d.input)
			if diff := cmp.Diff(d.expected, parsedData); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestDataProcessor(t *testing.T) {
	data := []struct {
		name string
		command []byte
		expected Result
	}{
		{"addition", []byte("CALC_1\n+\n3\n2"), Result{Id: "CALC_1", Value: 5}},
		{"subtraction", []byte("CALC_1\n-\n3\n2"), Result{Id: "CALC_1", Value: 1}},
		{"multiplication", []byte("CALC_1\n*\n3\n2"), Result{Id: "CALC_1", Value: 6}},
		{"division", []byte("CALC_1\n/\n3\n2"), Result{Id: "CALC_1", Value: 1}},
		{"division by 0", []byte("CALC_1\n/\n3\n0"), Result{Id: "", Value: 0}},
		{"arbitrary", []byte("CALC_1\n+\n3\na"), Result{Id: "", Value: 0}},
		{"arbitrary", []byte("CALC_1\n&\n3\na"), Result{Id: "", Value: 0}},
	}

	for _, d := range data {
		t.Run(d.name, func (t *testing.T) {
			in := make(chan []byte, 1)
			in <- d.command
			close(in)
			out := make(chan Result, 1)
			DataProcessor(in, out)
			if diff := cmp.Diff(d.expected, <-out); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestWriteData(t *testing.T) {
	expected := "CALC_1:6\n"
	buf := bytes.NewBuffer([]byte(""))
	in := make(chan Result, 1)
	in<- Result{Id: "CALC_1", Value: 6}
	close(in)
	WriteData(in, buf)
	msg, _ := buf.ReadString('\n')
	if diff := cmp.Diff(expected, msg); diff != "" {
		t.Error(diff)
	}
}

func TestNewController(t *testing.T) {
	data := []struct {
		name string
		expression string
		expectedCode int
		expectedMessage string
	}{
		{"case1", "CALC_1\n+\n3\n2", 202, "OK: 1"},
		{"case3", "CALC_1\n+\n3\n2", 202, "OK: 2"},
		{"case2", "CALC_1\n+\n3\n2", 503, "Too Busy: 1"},
		{"case3", "CALC_1\n+\n3\n2", 503, "Too Busy: 2"},
		{"case4", "bret", 400, "Bad Input"},
	}
	out := make(chan []byte, 2)
	server := httptest.NewServer(NewController(out))
	rc := RemoteCalc{
		CalcURL: server.URL,
		Client: server.Client(),
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T){

			res, err := rc.Calculate(d.expression)
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}

			if diff := cmp.Diff(res, d.expectedMessage); diff != "" {
				t.Error(diff)
			}
		})
	}
}
