package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {

	var ops uint64 = 0

	// 创建五十个协程
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1) // 原子级增值
				runtime.Gosched()         // 本协程让出CPU时间片段
			}
		}()
	}

	time.Sleep(time.Second)             // 休眠一秒
	opsFinal := atomic.LoadUint64(&ops) // 读取原子计数器
	fmt.Println("ops:", opsFinal)
}
