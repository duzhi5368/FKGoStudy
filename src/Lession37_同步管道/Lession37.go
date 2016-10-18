package main

import "fmt"
import "time"

func worker(done chan string) {
	fmt.Println("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("work done")
	time.Sleep(time.Second)
	done <- "ok"       // 通过这个赋值，导致下面的fmt.Println卡死等待
	fmt.Println("die") //	此处的die和主线程中的ok顺序是不稳定的
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	done := make(chan string, 1)
	go worker(done)
	fmt.Println(<-done) // 这里会无限卡死等待done管道中有数据，引发同步

	fmt.Println("===============")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs) // 这里的数据被两次管道传输
}
