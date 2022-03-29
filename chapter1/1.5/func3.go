package main

import "fmt"

func change(a, b int) (x, y int) {
	x = a + 100
	y = b + 100
	return
}

func main() {
	a := 1
	b := 2
	c, d := change(a, b)
	fmt.Println(c, d)
}
