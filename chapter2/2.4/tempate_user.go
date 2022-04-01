package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name string
	Sex  string
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("./template_user.tmpl")
	if err != nil {
		fmt.Println("err", err)
	}
	files.Execute(w, User{Name: "Jaluik", Sex: "male"})
}

func main() {
	http.HandleFunc("/", userHandler)
	http.ListenAndServe(":8081", nil)
}
