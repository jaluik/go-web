package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.MatchString("H(.*)d!", "Hello world!")
	if err != nil {
		return
	}
	fmt.Println("Matched: ", matched)
	reg, err := regexp.Compile(`^age is \d+`)
	if err != nil {
		return
	}
	matchString := reg.MatchString("age is 18")
	fmt.Println("Age Matched:", matchString)

	text1 := "Hello Gopher, Hello Go Web"
	reg1 := regexp.MustCompile(`\w+`)
	fmt.Println("Text1 match result:", reg1.FindAllString(text1, -1))
}
