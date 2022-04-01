package main

import "net/http"

type Refer struct {
	handler http.Handler
	refer   string
}

func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == this.refer {
		this.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is myHandler"))
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is hello handler"))
}

func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:   "www.jaluik.com",
	}
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", referer)
}
