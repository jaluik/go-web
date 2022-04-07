package main

import (
	"log"
	"os"
)

func main() {
	err := os.MkdirAll("test/test1/test3", 0777)
	if err != nil {
		log.Fatalln("创建文件失败：", err)
		return
	}
}
