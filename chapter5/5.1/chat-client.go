package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func Log(v ...string) {
	log.Println(v)
}

func Sender(conn *net.TCPConn) {
	defer conn.Close()
	sc := bufio.NewReader(os.Stdin)
	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			<-t.C
			_, err := conn.Write([]byte("1"))
			if err != nil {
				fmt.Println("发送心跳失败：", err)
				return
			}
		}
	}()
	name := ""
	fmt.Println("请输入聊天昵称")
	_, _ = fmt.Fscan(sc, &name)
	msg := ""
	buffer := make([]byte, 1024)
	_t := time.NewTicker(time.Second * 5)
	defer _t.Stop()
	go func() {
		<-_t.C
		fmt.Println("服务器出现故障")
		return
	}()

	for {
		go func() {
			for {
				n, err := conn.Read(buffer)
				if err != nil {
					fmt.Println("读取数据错误:", err)
					return
				}
				_t.Reset(time.Second * 5)
				if string(buffer[:n]) != "1" {
					fmt.Println(string(buffer[:n]))
				}
			}
		}()
		fmt.Fscan(sc, &msg)
		i := time.Now().Format("2006-01-02 15:04:05")
		_, _ = conn.Write([]byte(fmt.Sprintf("%s\n\t%s:%s", i, name, msg)))
	}

}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("解析域名错误:", err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal("链接服务器失败:", err)
		return
	}
	Log(conn.RemoteAddr().String(), "--链接服务器成功")
	Sender(conn)
	Log("结束")

}
