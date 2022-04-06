package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ArgLanguage struct {
	Java, Go string
}

type Programmer string

func (p Programmer) GetSkill(al ArgLanguage, skill *string) error {
	*skill = "Skill1: " + al.Java + ", Skill2: " + al.Go
	return nil
}

func main() {
	str := new(Programmer)
	err := rpc.Register(str)
	if err != nil {
		log.Fatalln(err)
		return
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("解析Tcp地址错误", err)
		return
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connect error:", err)
			continue
		}
		jsonrpc.ServeConn(conn)

	}

}
