package main

import "fmt"

func main() {
	x, y := 6, 8

	defer func(a int) {
		fmt.Printf("defer x, y = %d, %d\n", a, y)
	}(x)

	x += 10
	y += 100
	fmt.Printf("x, y = %d, %d\n", x, y)
}
