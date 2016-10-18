package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := "../../"
	// 调用系统函数
	// 该系统函数的第二个参数必须为指定参数和返回值的函数指针
	filepath.Walk(path, MyWalker)
}

// 该函数被会系统多次调用，fi是文件信息
func MyWalker(fn string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Walker Error: ", err)
		return nil
	}

	if fi.IsDir() {
		fmt.Println("Directory: ", fn)
	} else {
		fmt.Println("File: ", fn)
	}
	return nil
}
