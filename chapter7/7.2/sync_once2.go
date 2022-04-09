package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func fun1(c chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func fun2(c chan int, c2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-c
		if !ok {
			println("Not ok")
			break
		}
		c2 <- x * 2
	}
	once.Do(func() {
		close(c2)
	})
}

func main() {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	wg.Add(3)
	go fun1(c1)
	go fun2(c1, c2)
	go fun2(c1, c2)
	wg.Wait()
	for i := range c2 {
		fmt.Println(i)
	}

}
