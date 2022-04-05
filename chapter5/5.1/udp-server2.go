package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8013")
	if err != nil {
		fmt.Println("Resolve addr err:", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, resp, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println("收到客户端的消息：", string(buf[:n]))
		_, _ = conn.WriteToUDP([]byte("我是服务器"), resp)
	}
}
