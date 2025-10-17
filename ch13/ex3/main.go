package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Time struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	mux := http.NewServeMux()
	mux.Handle("GET /", ipTracker(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		switch f := r.Header.Get("accept"); f == "application/json" {
		case true:
			o := Time{
				DayOfWeek:  t.Weekday().String(),
				DayOfMonth: t.Day(),
				Month:      t.Month().String(),
				Year:       t.Year(),
				Hour:       t.Hour(),
				Minute:     t.Minute(),
				Second:     t.Second(),
			}
			out, err := json.Marshal(o)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Unexpected server error\n"))
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s\n", out)
		default:
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "The time is %s\n", t.Format(time.RFC3339))
		}
	})))
	s := http.Server{
		Addr:         ":8000",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func ipTracker(h http.Handler) http.Handler {
	options := &slog.HandlerOptions{Level: slog.LevelInfo}
	handler := slog.NewJSONHandler(os.Stderr, options)
	mySlog := slog.New(handler)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mySlog.Info("requester ip",
			"ip", r.RemoteAddr,
		)
		h.ServeHTTP(w, r)
	})
}
