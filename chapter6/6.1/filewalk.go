package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	err := filepath.Walk("./chapter1", func(path string, info fs.FileInfo, err error) error {
		fmt.Println("Scanned path: ", path)
		fmt.Printf("Name: %s, Size: %d\n", info.Name(), info.Size())
		return nil
	})
	if err != nil {
		return
	}
}
