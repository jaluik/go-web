package main

import "fmt"

func main() {
	var address = "Chengdu, China"
	ptr := &address
	fmt.Printf("prt type: %T\n", ptr)
	fmt.Printf("address: %p\n", ptr)

	value := *ptr
	fmt.Printf("value type: %T\n", value)
	fmt.Printf("value: %s\n", value)
}
