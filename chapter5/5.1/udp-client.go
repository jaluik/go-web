package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8012")
	if err != nil {
		fmt.Println("Dia error:", err)
		return
	}
	defer conn.Close()
	_, err = conn.Write([]byte("你好，这是UDP客户端发来消息"))
	if err != nil {
		fmt.Println("write error:", err)
		return
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	fmt.Println("服务器发来:", string(buf[:n]))
}
