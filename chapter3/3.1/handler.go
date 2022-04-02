package main

import (
	"fmt"
	"log"
	"net/http"
)

type handler1 struct {
}

func (h handler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hi, handle1")
}

type handler2 struct {
}

func (h handler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hi, handle2")
}
func main() {

	handler1 := handler1{}
	handler2 := handler2{}

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: nil,
	}
	http.Handle("/handle1", &handler1)
	http.Handle("/handle2", &handler2)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
