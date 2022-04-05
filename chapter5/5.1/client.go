package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Ussage: %s host:port", os.Args[0])
		return
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	validateError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	validateError(err)
	result, err := fullyRead(conn)
	validateError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func validateError(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal Error is error:", err)
		os.Exit(1)
	}
}

func fullyRead(conn net.Conn) ([]byte, error) {
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return nil, err
		}
		result.Write(buf[0:n])
		if err != nil {
			if err != io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
