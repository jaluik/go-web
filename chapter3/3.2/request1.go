package main

import "net/url"

func main() {
	addr := "http://localhost:8080/article?id=1"

	parse, _ := url.Parse(addr)
	println(parse.Host)
	println(parse.Path)
	println(parse.User)
	println(parse.RawQuery)
	println(parse.RequestURI())
}
