package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}
	fmt.Println("Locked")
	mutex.Lock()
	for i := 0; i < 5; i++ {
		wait.Add(1)
		go func(i int) {
			fmt.Println("Not Locked: ", i)
			mutex.Lock()
			fmt.Println("Lock:", i)
			time.Sleep(time.Second)
			fmt.Println("Unlocked:", i)
			mutex.Unlock()
			defer wait.Done()
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	mutex.Unlock()
	wait.Wait()

}
