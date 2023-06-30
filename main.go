package main

import (
	"time"
	"net/http"
	"fmt"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Jiro!\n"))
}

func main() {

	http.HandleFunc("/health", HealthCheck)

	s := http.Server{
		Addr: ":8080",
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout: 90 * time.Second,
	}

	fmt.Println("Server Start")

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}