package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(file)
		return
	}
	defer file.Close()
	_, err = file.WriteString("Hello Go!")
	if err != nil {
		log.Fatal(err)
		return
	}
	_, _ = file.WriteAt([]byte("编程使我快乐"), 66)

}
