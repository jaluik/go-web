package main

import "fmt"

func fibonacci(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}
	close(c)
}

func main() {
	ch := make(chan int, 6)
	go fibonacci(cap(ch), ch)
	for v := range ch {
		fmt.Println(v)
	}
}
