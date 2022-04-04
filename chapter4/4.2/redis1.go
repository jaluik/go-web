package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	dial, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Can connect redis, err:", err)
		return
	}
	_, err = dial.Do("Set", "username", "jack")
	if err != nil {
		return
	}

	res, err := redis.String(dial.Do("Get", "username"))
	if err != nil {
		return
	}
	fmt.Println(res)
	defer func(dial redis.Conn) {
		err := dial.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(dial)

}
