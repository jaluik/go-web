package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("LICENSE")
	if err != nil {
		log.Fatal("打开文件失败了: ", err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("读取文件失败了：", err)
		return
	}
	fmt.Printf("%v", string(content))
}
