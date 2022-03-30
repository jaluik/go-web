package main

import "fmt"

func Len(array interface{}) int {
	var length int
	if array == nil {
		length = 0
	}
	switch array.(type) {
	case []int:
		length = len(array.([]int))
	case []string:
		length = len(array.([]string))
	default:
		length = 0
	}
	return length

}

func main() {
	fmt.Println(Len([]int{3, 4, 5}))
}
