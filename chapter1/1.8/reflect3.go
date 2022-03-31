package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "jaluik"

	//需要传入指针
	v := reflect.ValueOf(&name)
	fmt.Println("可写性为： ", v.CanSet())

	v1 := v.Elem()
	fmt.Println("可写性为： ", v1.CanSet())
}
