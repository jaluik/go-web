package main

import "fmt"

func main() {
	var a interface{} = func(a int) string {
		return fmt.Sprintf("d: %d", a)
	}
	switch b := a.(type) {
	case nil:
		fmt.Println("Nil")
	case *int:
		fmt.Println(*b)
	case func(int) string:
		fmt.Println(b(66))
	default:
		fmt.Println("no match")
	}

}
