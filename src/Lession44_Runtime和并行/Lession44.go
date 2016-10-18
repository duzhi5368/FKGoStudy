package main

import (
	"fmt"
	"runtime"
	"time"
)

var quit chan int

func foo(id int) {
	fmt.Println(id)
	time.Sleep(time.Second) // 停顿一秒
	quit <- 0               // 发消息：我执行完啦！
}

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("当前CPU有核心数：", runtime.NumCPU())
	runtime.GOMAXPROCS(4) // 设置GO可用的核心数
	begin := time.Now()
	count := 1000
	quit = make(chan int, count) // 缓冲1000个数据

	for i := 0; i < count; i++ { // 开1000个goroutine
		go foo(i)
	}

	for i := 0; i < count; i++ { // 等待所有完成消息发送完毕。
		<-quit
	}
	end := time.Now()
	fmt.Println("累计消耗时间：",
		end.Second()-begin.Second(), "秒",
		end.Nanosecond()-begin.Nanosecond(), "纳秒")

	// 如果我们runtime.GOMAXPROCS(1)设置为1，则不会有任何输出
	// 因为其实当前GO协程是在进行单CPU的切片模拟多线程而已。
	// 如果我们runtime.GOMAXPROCS(4)，则可以正常输出五次world
	// 因为此时我们确实多CPU并行

	// 我们要明确要求CPU并行，则需要
	//	1：可以runtime.GOMAXPROCS(x)指定多CPU
	//  2：可以runtime.Gosched显式调用表示让出CPU切片
	go say("world") //开一个新的Goroutines执行
	for {
	}
}
