package main

import (
	"fmt"
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
	mux.HandleFunc("GET /time", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("The time is %s\n", time.Now().Format(time.RFC3339))))
	})
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
