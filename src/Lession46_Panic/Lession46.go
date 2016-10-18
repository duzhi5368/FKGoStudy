package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {
	_, err := os.Create(getCurrentDirectory() + "/test.txt")
	if err != nil {
		panic(err)
	}
	// 这里会宕掉，并控制台输出堆栈
	panic("a problem")
}
