package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	mux := http.NewServeMux()
	mux.Handle("GET /", ipTracker(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("The time is %s\n", time.Now().Format(time.RFC3339))))
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
