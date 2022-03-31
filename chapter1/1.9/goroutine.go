package main

import (
	"fmt"
	"time"
)

func Hello() {
	fmt.Println("Hello world")
}

func main() {
	go Hello()
	time.Sleep(1 * time.Second)
	fmt.Println("End")
}
