package main

import "fmt"

func main() {
	var score int = 10
	var name string = "Barry"
	// %p 取指针的值
	fmt.Printf("%p %p", &score, &name)
}
