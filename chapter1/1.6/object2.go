package main

import (
	"fmt"
	"go-web/chapter1/1.6/person"
)

func main() {
	s := person.Student{}
	s.SetName("张三")
	fmt.Println(s.GetName())
}
