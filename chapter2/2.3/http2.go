package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection: %s", r.Proto)
	w.Write([]byte("This is a message from http2"))
}

func main() {
	serv := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(handler)}
	log.Printf("Serving on https://localhost:8080")
	//需要执行 openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt 生成密匙对
	log.Fatal(serv.ListenAndServeTLS("./server.crt", "./server.key"))
}
