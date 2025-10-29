package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	mux := http.NewServeMux()
	mux.Handle("GET /", timeoutMiddlewareFactory(100)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wait := rand.Intn(200)
		ctx := r.Context()
		var message string
		select {
		case <-time.After(time.Duration(wait) * time.Millisecond):
			w.WriteHeader(http.StatusOK)
			message = "Done!"
		case <-ctx.Done():
			w.WriteHeader(http.StatusInternalServerError)
			message = "Too slow!"
		}
		w.Write([]byte(message))
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

func timeoutMiddlewareFactory(t int) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, cancelFunc := context.WithTimeout(ctx, time.Duration(t)*time.Millisecond)
			defer cancelFunc()
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}
