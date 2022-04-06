package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type Args struct {
	X, Y int
}

type Algorithm int

func (algorithm Algorithm) Sum(args *Args, reply *int) error {
	*reply = args.X + args.Y
	fmt.Println("Exec Sum: ", *reply)
	return nil
}

func main() {
	algorithm := new(Algorithm)
	fmt.Println("Algorithm Start:", algorithm)
	err := rpc.Register(algorithm)
	if err != nil {
		fmt.Println("注册失败了：", err)
		return
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("err========", err)
	}
}
