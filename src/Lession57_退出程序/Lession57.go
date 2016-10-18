package main

import "fmt"
import "os"

func main() {

	// 这个defer将永远无法被执行到
	defer fmt.Println("!")

	// 退出，返回值3。返回值只有0才认为是正常退出
	os.Exit(3)
}
