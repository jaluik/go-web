package main

import (
	"log"
	"os"
)

func main() {
	err := os.Mkdir("test", 0777)
	if err != nil {
		log.Fatalln("创建文件失败：", err)
		return
	}
}
