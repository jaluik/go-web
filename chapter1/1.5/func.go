package main

import "fmt"

func main() {
	array := []int{6, 8, 10, 3}
	ret := min(array)
	fmt.Printf("最小值： %d", ret)

}

func min(arr []int) (min int) {
	min = arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return
}
