package main

import "fmt"

func main() {
	var array [6]int
	for i := 0; i < 6; i++ {
		array[i] = i + 66
	}
	for j := 0; j < 6; j++ {
		fmt.Printf("Array[ %d ]  = %d\n", j, array[j])
	}
}
