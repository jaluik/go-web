package main

import "fmt"

func compute(x, y int) (int, int) {
	return x + y, x * y

}

func main() {
	a, b := compute(6, 10)
	fmt.Println(a, b)
}
