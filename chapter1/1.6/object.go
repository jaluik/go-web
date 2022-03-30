package main

import "fmt"

type Triangle struct {
	Bottom float32
	Height float32
}

func (t *Triangle) Area() float32 {
	return (t.Height * t.Bottom) / 2
}

func main() {
	r := Triangle{3, 4}
	fmt.Println(r.Area())
}
