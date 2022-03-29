package main

import "fmt"

func visitPrint(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func main() {
	list := []int{6, 8, 10}
	visitPrint(list, func(i int) {
		fmt.Println(i)
	})
}
