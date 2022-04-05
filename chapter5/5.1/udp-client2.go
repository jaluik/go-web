package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8013")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()
	go func() {
		str := make([]byte, 1024)
		for {

			n, err2 := os.Stdin.Read(str)
			if err2 != nil {
				fmt.Println("read stdin err:", err)
				return
			}
			_, _ = conn.Write(str[0:n])
		}
	}()
	buf := make([]byte, 1024)
	for {
		n, err2 := conn.Read(buf)
		if err2 != nil {
			fmt.Println("read udp err:", err2)
			return
		}
		fmt.Println("收到服务器的消息：", string(buf[:n]))
	}
}
