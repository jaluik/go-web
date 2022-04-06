package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Send struct {
	Java, Go string
}

func main() {
	sender := Send{"java", "go"}
	dial, err := jsonrpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("连接服务器失败：", err)
		return
	}
	defer dial.Close()
	var receiver string
	err = dial.Call("Programmer.GetSkill", sender, &receiver)
	if err != nil {
		log.Fatal("计算失败了：", err)
		return
	}
	fmt.Printf("得到的内容是：%s\n", receiver)

}
