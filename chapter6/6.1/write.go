package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("write1.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	content := []byte("你好， web")
	_, err = file.Write(content)
	if err != nil {
		fmt.Println("文件写入错误", err)
		return
	}
	fmt.Println("文件写入成功！")

}
