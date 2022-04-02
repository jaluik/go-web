package main

import (
	"fmt"
	"log"
	"net/http"
)

type welcomeHandler struct {
}

func (t welcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome handler")
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi handler")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", hiHandler)
	mux.Handle("/welcome", welcomeHandler{})
	server := http.Server{
		Addr:    "0.0.0.0:80",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
