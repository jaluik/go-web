package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	p := &redis.Pool{
		MaxIdle:     16,
		MaxActive:   1024,
		IdleTimeout: 30,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	conn := p.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	_, err := conn.Do("Set", "phone", "18580001234")
	if err != nil {
		return
	}
	res1, err := redis.String(conn.Do("Get", "username"))
	if err != nil {
		log.Fatal("res1 err:", err)
		return
	}
	fmt.Println(res1)
	res2, err := redis.String(conn.Do("Get", "phone"))
	if err != nil {
		log.Fatal("res2 err:", err)
		return
	}
	fmt.Println(res2)
}
