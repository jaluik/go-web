package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func CopyHandle(method, urlValue, data string) {
	client := &http.Client{}
	var req *http.Request
	if data == "" {
		req, _ = http.NewRequest(method, urlValue, nil)
	} else {
		req, _ = http.NewRequest(method, urlValue, strings.NewReader(data))
	}
	cookie := &http.Cookie{
		Name:     "X-Xsrftoken",
		Value:    "xxxxxxx",
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	req.Header.Add("X-Xsrftoken", "xxxxxx")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	b, _ := ioutil.ReadAll(resp.Body)
	_, err = fmt.Println(string(b))
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	CopyHandle("GET", "http://www.baidu.com", "")
}
