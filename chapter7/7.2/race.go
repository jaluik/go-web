package main

import "fmt"

func main() {
	c := make(chan bool)
	m := make(map[string]string)

	go func() {
		m["a"] = "one"
		c <- true
	}()
	m["a"] = "two"
	<-c
	for key, value := range m {
		fmt.Println(key, value)
	}
}
