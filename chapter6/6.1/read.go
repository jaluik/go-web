package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("LICENSE")
	if err != nil {
		log.Fatal("打开文件失败:", err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	reader := bufio.NewReader(file)
	for {
		readLine, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Print(readLine)
	}
}
