package main

import "fmt"

func main() {
	defer fun1()
	defer fun2()
	defer fun3()
}

func fun1() {
	fmt.Println("func1 is called")
}
func fun2() {
	fmt.Println("func2 is called")
}
func fun3() {
	fmt.Println("func3 is called")
}
