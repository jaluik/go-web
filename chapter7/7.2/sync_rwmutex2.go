package main

import (
	"fmt"
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)
	go Reading(1)
	go Reading(2)
	time.Sleep(2 * time.Second)

}

func Reading(i int) {
	fmt.Println(i, "读开始了")
	m.RLock()
	fmt.Println(i, "正在读取")
	time.Sleep(1 * time.Second)
	m.RUnlock()
	fmt.Println(i, "读取结束了")
}
