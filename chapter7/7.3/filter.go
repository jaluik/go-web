package main

import "fmt"

func main() {
	const max = 100
	numbers := IntegerGenerator()
	number := <-numbers

	for number <= max {
		fmt.Println(number)
		numbers = Filter(numbers, number)
		number = <-numbers
	}

}

func IntegerGenerator() chan int {
	ch := make(chan int)

	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func Filter(ch chan int, n int) chan int {
	c1 := make(chan int)
	go func() {
		for {
			v := <-ch
			if v%n != 0 {
				c1 <- v
			}

		}
	}()

	return c1
}
