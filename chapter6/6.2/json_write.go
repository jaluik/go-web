package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	UserName string
	NickName string `json:"nickName"`
	Email    string
}

func main() {
	user := &User{
		UserName: "jaluik",
		NickName: "no",
		Email:    "cqzhengmouren@gmail.com",
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Json string is: ", string(bytes))
	file, err := os.Create("test.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	_, err = file.Write(bytes)
	if err != nil {
		fmt.Println("写入json错误：", err)
		return
	}
}
