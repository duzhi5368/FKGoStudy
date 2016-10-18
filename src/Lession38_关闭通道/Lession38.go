package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool) // 通知关闭的信号通道

	// 开了个协程，在等待通道
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 主进程插入了三个通道数据
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // 主进程主动关闭通道，如果不关闭，则协程还会在等待主线程发送通道，结果会互相等待，死锁
	fmt.Println("sent all jobs")

	// 主进程挂起，等待协程通知关闭
	fmt.Println(<-done)
}
