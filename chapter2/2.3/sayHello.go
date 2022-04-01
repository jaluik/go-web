package main

import "net/http"

func SayHello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/", SayHello)
	http.ListenAndServe(":8080", nil)
}
