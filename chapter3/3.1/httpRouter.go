package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "http router index")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
