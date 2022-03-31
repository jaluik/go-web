package testexample

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	array := []int{6, 8, 10}
	min := Min(array)
	fmt.Println(min)
}
