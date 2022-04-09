package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	calledFn := func() {
		fmt.Println("This will be called only once")
	}
	ch := make(chan bool)
	for i := 0; i < 3; i++ {
		go func() {
			once.Do(calledFn)
			ch <- true
		}()
	}
	for i := 0; i < 3; i++ {
		<-ch
	}
	defer close(ch)
}
