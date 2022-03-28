package main

import "fmt"

func main() {
	var local1, local2, local3 int

	local1 = 8
	local2 = 10
	local3 = local1 * local2

	fmt.Printf("local1 is %d, local2 is %d, local3 is %d", local1, local2, local3)
}
