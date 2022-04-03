package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Greeting struct {
	Message string `json:"message"`
}

func Hello(w http.ResponseWriter, _ *http.Request) {
	greeting := Greeting{
		"This is a message",
	}
	message, _ := json.Marshal(greeting)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
