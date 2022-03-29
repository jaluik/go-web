package main

import "fmt"

func exchange(a, b *int) {
	tem := *a
	*a = *b
	*b = tem

}

func main() {
	a, b := 1, 3

	fmt.Printf("a is %d, b is %d\n", a, b)
	exchange(&a, &b)
	fmt.Printf("a is %d, b is %d", a, b)
}
