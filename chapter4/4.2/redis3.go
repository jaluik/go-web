package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal("Init failed, err:", err)
		return
	}
	defer func(conn redis.Conn) {
		_ = conn.Close()
	}(conn)

	_ = conn.Send("MULTI")
	_ = conn.Send("INCR", "foo")
	_ = conn.Send("INCR", "bar")
	res, err := conn.Do("EXEC")
	if err != nil {
		return
	}
	fmt.Println(res)
}
