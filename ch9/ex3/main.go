package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var ErrInvalidID = errors.New("invalid id")

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}

		myError := ValidateEmployee(emp)
		if myError.Errors != nil {
			for _, err := range myError.Unwrap() {
				var emptyFieldErr EmptyFieldErr
				if err != nil {
					if errors.Is(err, ErrInvalidID) {
						fmt.Printf("record %d: %+v error: invalid ID: %s\n", count, emp, emp.ID)
					} else if errors.As(err, &emptyFieldErr) {
						fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
					} else {
						fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
					}
				}
			}
			continue
		}

		fmt.Printf("record %d: %+v\n", count, emp)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

type MyError struct {
	Errors []error
}

func (err MyError) Error() string {
	return errors.Join(err.Errors...).Error()
}

func (err MyError) Unwrap() []error {
	return err.Errors
}

type EmptyFieldErr struct {
	field string
}

func (err EmptyFieldErr) Error() string {
	return fmt.Sprintf("empty field: %s", err.field)
}

var (
	validID = regexp.MustCompile(`\W{4}-\d{3}`)
)

func ValidateEmployee(e Employee) MyError {
	myError := MyError{
		Errors: make([]error, 4),
	}
	if len(e.ID) == 0 {
		myError.Errors = append(myError.Errors, EmptyFieldErr{field: "ID"})
	}
	if !validID.MatchString(e.ID) {
		myError.Errors = append(myError.Errors, ErrInvalidID)
	}
	if len(e.FirstName) == 0 {
		myError.Errors = append(myError.Errors, EmptyFieldErr{field: "FirstName"})
	}
	if len(e.LastName) == 0 {
		myError.Errors = append(myError.Errors, EmptyFieldErr{field: "LastName"})
	}
	if len(e.Title) == 0 {
		myError.Errors = append(myError.Errors, EmptyFieldErr{field: "Title"})
	}

	return myError
}
