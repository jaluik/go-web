package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name interface{} = "jaluik"

	fmt.Printf("原始接口类型变量的类型是%T, 值为%v \n", name, name)

	t := reflect.TypeOf(name)
	v := reflect.ValueOf(name)

	fmt.Printf("从接口类型到反射对象，Type对象的类型是%T \n", t)
	fmt.Printf("从接口类型到反射对象，Value对象的类型是%T \n", v)

	i := v.Interface()
	fmt.Printf("原始接口类型变量的类型是%T, 值为%v \n", i, i)
}
