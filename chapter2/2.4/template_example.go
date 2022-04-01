package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func templateHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template_example.tmpl")
	if err != nil {
		fmt.Println("err", err)
	}
	name := "jaluik"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", templateHandler)
	http.ListenAndServe(":8080", nil)
}
