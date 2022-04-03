package main

import (
	"fmt"
	"log"
	"net/http"
)

type welcomHandler struct {
	Language string
}

func (w welcomHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "language is %s", w.Language)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/cn", welcomHandler{"Welcome to Chinese"})
	mux.Handle("/en", welcomHandler{"Welcome to English"})
	server := http.Server{Addr: ":8080", Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
