package main

import "fmt"

var name string = "go"

func myFunc() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myFunc中的name是%s\n", name)

	return name
}

func main() {
	newName := myFunc()
	fmt.Printf("main函数里面的name是%s\n", name)
	fmt.Printf("main函数中返回的name是%s\n", newName)
}
