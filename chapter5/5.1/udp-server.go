package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8012")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Listen UDP err:", err)
		return
	}
	defer func(conn *net.UDPConn) {
		_ = conn.Close()
	}(conn)

	buf := make([]byte, 1024)
	n, readAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		return
	}
	fmt.Println("客户端发送： ", string(buf[0:n]))
	_, err = conn.WriteToUDP([]byte("你好客户端，我是服务器"), readAddr)
	if err != nil {
		fmt.Println("write udp err:", err)
		return
	}
}
