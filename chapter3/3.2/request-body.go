package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getBody(w http.ResponseWriter, r *http.Request) {
	contentLen := r.ContentLength
	body := make([]byte, contentLen)
	_, err := r.Body.Read(body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(w, string(body))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/getBody", getBody)
	http.ListenAndServe(":8080", nil)
}
