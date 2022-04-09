package main

import "fmt"

func Sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{6, 7, 8, 8, 3, 5, -1, 8}
	ch := make(chan int)
	go Sum(s[len(s)/2:], ch)
	go Sum(s[:len(s)/2], ch)
	sum1, sum2 := <-ch, <-ch
	fmt.Printf("%d + %d = %d \n", sum1, sum2, sum1+sum2)
}
