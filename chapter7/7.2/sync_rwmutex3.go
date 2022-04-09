package main

import (
	"sync"
	"time"
)

var m1 *sync.RWMutex

func Read(i int) {
	println(i, "读取开始了")
	m1.RLock()
	println(i, "正在读取")
	time.Sleep(1 * time.Second)
	m1.RUnlock()
	println(i, "读取结束了")
}

func write(i int) {
	println(i, "写入开始了")
	m1.Lock()
	println(i, "正在写入")
	time.Sleep(1 * time.Second)
	m1.Unlock()
	println(i, "写入结束了")
}

func main() {
	m1 = new(sync.RWMutex)
	go write(1)
	go Read(2)
	go write(3)
	time.Sleep(3 * time.Second)
}
