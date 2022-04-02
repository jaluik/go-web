package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not bad")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hi)
	server := http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
