package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

type logKey int

const (
	_ logKey = iota
	key
)

type ResponseError struct {
	StatusCode int    `json:"status_code"`
	Message      string `json:"message"`
}

func main() {
}

func storeLogLevel(level Level, ctx context.Context) context.Context {
	return context.WithValue(ctx, key, level)
}

func extractLogLevel(ctx context.Context) Level {
	return ctx.Value(key).(Level)
}

func middleLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logLevel := r.URL.Query().Get("log_level")
		if logLevel != "debug" && logLevel != "info" {
			w.WriteHeader(http.StatusBadRequest)
			o := ResponseError{
				StatusCode: 400,
				Message: "Expected log_level of info or debug",
			}
			out, err := json.Marshal(o)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Unexpected server error\n"))
			}
			fmt.Fprintf(w, "%s\n", out)
		}

		h.ServeHTTP(w, r)
	})
}

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	// TODO: get a logging level out of the context and assign it to inLevel
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}

	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}
