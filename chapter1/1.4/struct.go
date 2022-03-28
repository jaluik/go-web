package main

import "fmt"

type Book struct {
	title, author, subject, press string
}

func main() {
	fmt.Println(Book{"名字", "jaluik", "go", "no publish"})
	fmt.Println(Book{title: "名字1", author: "jaluik"})
}
