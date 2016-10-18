package main

// 我们将使用协程的内置同步Channels来实现多协程共享互斥，可替代mutex
// 该方法和GO的共享内存机制比较类似

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 我们这里一个状态，仅被一个协程管理,
// 这样可以保证数据永远不会被并发访问和损坏.
// 为了对这个状态进行读写，其他状态必须向该协程请求。
// 下面就是请求消息
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var ops int64 = 0 // 执行的操作数，原子计数

	// 读写管道
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// 该协程负责管理状态，也就是所谓的“状态协程”
	// 它会不停的尝试从通道获取请求并处理，做出响应
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 这里启动了100个协程
	// 它们每次启动读取都需要向状态协程提交请求
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 同样，创建10个写入协程
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 主线休眠1秒，等待协程自己处理
	time.Sleep(time.Second)

	// 输出执行次数
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}
