package main

import "os"
import "strings"
import "fmt"

func main() {
	// 设置获取环境变量测试
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	fmt.Println()

	// 输出当前系统环境变量
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], " = ", pair[1])
	}
}
