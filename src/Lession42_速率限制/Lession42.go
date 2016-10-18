package main

import "time"
import "fmt"

func main() {
	// 我们提交了一些工作请求
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter通道每200毫秒会接受一次值
	limiter := time.Tick(time.Millisecond * 200)

	// 我们通过从limiter通道中取值，来限制200毫秒才能进行一次消息发送
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("===================")

	// 我们许可有三个突发事件通道队列
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每200毫秒我们添加一个突发事件到通道队列中
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 我们当前模拟5个请求。前三个会在我们的紧急突发通道列表中
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests) // 关闭通道
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
