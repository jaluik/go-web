package main

import (
	"fmt"
	"net"
	"time"
)

type HeartBeat struct {
	endTime int64
}

var ConnSlice map[net.Conn]*HeartBeat

func handleConn(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if ConnSlice[conn].endTime > time.Now().Unix() {
			ConnSlice[conn].endTime = time.Now().Add(time.Second * 5).Unix()
		} else {
			fmt.Println("长时间未发送消息断开链接：")
			return
		}
		if err != nil {
			return
		}
		// heartbeat test
		if string(buf[:n]) == "1" {
			_, _ = conn.Write([]byte("1"))
			continue
		}
		for c, heartBeat := range ConnSlice {
			if c == conn {
				continue
			}
			if heartBeat.endTime < time.Now().Unix() {
				delete(ConnSlice, c)
				c.Close()
				fmt.Println("删除连接:", c.RemoteAddr())
				fmt.Println("现存链接", ConnSlice)
			}
			_, err = c.Write(buf[:n])
			if err != nil {
				fmt.Println("Write error:", err)
				return
			}
		}
	}
}

func main() {

	ConnSlice = map[net.Conn]*HeartBeat{}
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("服务器启动失败", err)
		return
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Connect error:", err)
			return
		}
		ConnSlice[conn] = &HeartBeat{
			endTime: time.Now().Add(time.Second * 5).Unix(),
		}
		go handleConn(conn)

	}
}
