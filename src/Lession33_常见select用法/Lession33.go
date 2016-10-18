package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个chan
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2 * 1e9) // 休眠2秒
		timeout <- true     // 实现TimeOut机制
	}()
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("ch pop one element")
	case <-timeout:
		fmt.Println("timeout!")
	}

	fmt.Println("-------------")
	select {
	case ch <- 2:
		// 当ch满了，此时向ch中插入2会失败，则会执行default
		// 此时会实现对channel是否已满的检查
	default:
		fmt.Println("channel is full !")
	}
}
