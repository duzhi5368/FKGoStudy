package main

import (
	"fmt"
	"time"
)

// c 和 quit 都是 int类型的Channel
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// select 和 switch非常相似，但是select强制要求其case必须是I/O操作
		select {
		case c <- x: // 主线程没有退出信号，则死循环尝试将 x填充到c Channal中,直到填充了10次
			x, y = y, x+y
		case <-quit: // 主线程发现quit channel有信号，则直接退出
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// 创建了一个协程，让它去输出 c channel的数据10次
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		// 输出结束，给quit发个信号
		quit <- 0
	}()
	fibonacci(c, quit)

	fmt.Println("==========================")
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// 注意for循环的次数用来检查接收次数
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)

		}
	}
}
