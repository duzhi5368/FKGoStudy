package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings" // 字符串操作
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
		//filepath.Walk(fn, MyWalker)
	} else {
		if strings.Contains(fn, ".exe") && !strings.Contains(fn, "Tools") {
			DeleteFile(fn)
		}
	}
	return nil
}

// 删除文件
func DeleteFile(filepath string) {
	err := os.Remove(filepath)
	if err != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Println("file remove Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	} else {
		fmt.Println("删除文件 ：%s", filepath)
	}
}
