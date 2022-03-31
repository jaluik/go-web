package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Go web learning"
	fmt.Println("真实的value值为", name)

	v := reflect.ValueOf(&name)
	v1 := v.Elem()
	v1.SetString("Changed name")
	fmt.Println("更新过的value值为", name)
}
