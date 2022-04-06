package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

type ArgsTwo struct {
	X, Y int
}

func main() {
	dial, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln("请求时发生错误: ", err)
		return
	}
	defer func(dial *rpc.Client) {
		_ = dial.Close()
	}(dial)
	i1, _ := strconv.Atoi(os.Args[1])
	i2, _ := strconv.Atoi(os.Args[2])
	args := ArgsTwo{i1, i2}
	var reply int
	err = dial.Call("Algorithm.Sum", args, &reply)
	if err != nil {
		log.Fatal("计算错误：", err)
		return
	}
	fmt.Printf("%d + %d = %d \n", args.X, args.Y, reply)
}
