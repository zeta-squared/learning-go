package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a type called `Level` whose underlying type is string
type Level string

// Define two constants of type `Level`, `Debug` and `Info`, and set them
// to "debug" and "info", respectively.
const (
	Debug Level = "debug"
	Info  Level = "info"
)

type logKey int

const (
	_ logKey = iota
	key
)

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	mux := http.NewServeMux()

	mux.Handle("/info", middleLog(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Log(r.Context(), Info, "[INFO]: Hello, World!")
		o := Response{
			StatusCode: 200,
			Message:    "Message logged",
		}
		out, err := json.Marshal(o)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unexpected server error\n"))
		}
		fmt.Fprintf(w, "%s\n", out)
	})))

	mux.Handle("/debug", middleLog(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Log(r.Context(), Debug, "[DEBUG]: Hello, World!")
		o := Response{
			StatusCode: 200,
			Message:    "Message logged",
		}
		out, err := json.Marshal(o)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unexpected server error\n"))
		}
		fmt.Fprintf(w, "%s\n", out)
	})))

	s := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

// Create functions to store the log level in the context and to extract it.
func storeLogLevel(level Level, ctx context.Context) context.Context {
	return context.WithValue(ctx, key, level)
}

func extractLogLevel(ctx context.Context) (Level, bool) {
	level, ok := ctx.Value(key).(Level)
	return level, ok
}

// Create a middleware function to get the logging level from a query paramter
// called `log_level`. The valid values for `log_level` are `debug` and `info`.
func middleLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logLevel := Level(r.URL.Query().Get("log_level"))
		if logLevel != Debug && logLevel != Info {
			w.WriteHeader(http.StatusBadRequest)
			o := Response{
				StatusCode: 400,
				Message:    "Expected log_level of info or debug",
			}
			out, err := json.Marshal(o)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Unexpected server error\n"))
			}
			fmt.Fprintf(w, "%s\n", out)
		}
		r = r.WithContext(storeLogLevel(logLevel, r.Context()))

		h.ServeHTTP(w, r)
	})
}

// Fill in the `TODO` in `Log` to properly extract the log level from the context.
// If the log level is not assigned or is not a valid value, nothing should be printed.
func Log(ctx context.Context, level Level, message string) {
	inLevel, ok := extractLogLevel(ctx)
	if !ok {
		return
	}

	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}

	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}
