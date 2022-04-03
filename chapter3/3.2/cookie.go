package main

import (
	"log"
	"net/http"
)

func handleCookie(w http.ResponseWriter, _ *http.Request) {
	cookie := &http.Cookie{
		Name:   "test_cookie",
		Value:  "cookie_value",
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/cookie",
	}
	http.SetCookie(w, cookie)
	_, err := w.Write([]byte("Cookie 设置成功"))
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	http.HandleFunc("/cookie", handleCookie)
	http.ListenAndServe(":8080", nil)
}
