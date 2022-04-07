package main

import (
	"log"
	"os"
)

func main() {
	err := os.Mkdir("old_name", 0777)
	if err != nil {
		log.Fatalln(err)
		return
	}
	oldName := "old_name"
	newName := "new_name"

	err = os.Rename(oldName, newName)
	if err != nil {
		log.Fatal("重命名失败：", err)
		return
	}
}
