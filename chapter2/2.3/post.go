package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://www.shirdon.com/comment/add"
	body := "{\"comment\": 2, \"articalId\":3, \"comment\": \"😥️这不是一条评论\"}"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("Error", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(string(b))
}
