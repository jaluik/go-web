package main

import (
	"fmt"
	"log"
	"net"
)

func Server() {
	l, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		log.Fatal(err)
	}
	defer func(l net.Listener) {
		_ = l.Close()
	}(l)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("访问客户端信息： conn=%v, 客户端 ip=%v \n", conn, conn.RemoteAddr().String())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	for {
		fmt.Printf("服务器在等待客户端 %s 发送消息\n", conn.RemoteAddr().String())
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err: ", err)
			break
		}
		fmt.Println(string(buf[0:n]))
	}

}

func main() {
	Server()
}
