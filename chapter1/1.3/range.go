package main

import "fmt"

func main() {
	//for key, value := range []int{0, 1, -1, -2} {
	//	fmt.Printf("key: %d, value: %d \n", key, value)
	//}
	//str := "hi 加油"
	//for key, value := range str {
	//	fmt.Printf("key: %d, value is %q \n", key, value)
	//
	//}
	//m := map[string]int{
	//	"go":  100,
	//	"web": 1000,
	//}
	//for key, value := range m {
	//	fmt.Printf("key is %v, value is %d", key, value)
	//}
	c := make(chan int)
	go func() {
		c <- 7
		c <- 8
		c <- 9
		close(c)
	}()
	for val := range c {
		fmt.Printf("value is %d \n", val)

	}
}
