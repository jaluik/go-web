package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	press   string
}

func printBook(book Books) {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book suject: %s\n", book.subject)
	fmt.Printf("Book press: %s\n", book.press)
}
func printPointerBook(book *Books) {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book suject: %s\n", book.subject)
	fmt.Printf("Book press: %s\n", book.press)
}

func main() {
	bookGo := Books{}
	bookGo.title = "Go Web"
	bookGo.author = "jaluik"
	bookGo.subject = "编程相关"
	bookGo.press = "xxx出版社"
	printBook(bookGo)
	printPointerBook(&bookGo)
}
