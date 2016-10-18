package main

import "time"
import "fmt"

func main() {
	ticker := time.NewTicker(time.Millisecond * 500) // 500毫秒一次触发
	go func() {
		for t := range ticker.C { // 循环触发
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Millisecond * 1600) // 1600毫秒后，关闭触发器
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
