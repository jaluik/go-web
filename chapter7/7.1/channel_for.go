package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	go func() {
		for i := 0; i <= 3; i++ {
			channel <- i
			time.Sleep(time.Second)
		}
		close(channel)
	}()
	for receive := range channel {
		fmt.Println(receive)
	}
}
