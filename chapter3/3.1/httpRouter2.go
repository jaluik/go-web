package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type HostMap map[string]http.Handler

func (receiver HostMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := receiver[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", 403)
	}
}

func main() {
	userRouter := httprouter.New()
	userRouter.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("sub1"))
	})

	dataRouter := httprouter.New()
	dataRouter.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("sub2"))
	})
	hs := make(HostMap)
	hs["sub1.localhost:8080"] = userRouter
	hs["sub2.localhost:8080"] = dataRouter
	log.Fatal(http.ListenAndServe(":8080", hs))
}
