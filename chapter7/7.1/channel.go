package main

import "fmt"

func main() {
	channel := make(chan string)
	go func() {
		fmt.Println("开始goroutine")
		channel <- "signal"
		fmt.Println("结束goroutine")
	}()
	fmt.Println("等待goroutine")
	<-channel
	fmt.Println("完成")
}
