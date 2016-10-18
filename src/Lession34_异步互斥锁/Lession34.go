package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// 类内部配有异步互斥锁
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock() // 如果去除这两个lock行为，将会报错提示写入冲突
	c.v[key]++
	c.mux.Unlock() // 如果去除这两个lock行为，将会报错提示写入冲突
}
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey") // 开启1000个协程，进行inc叠加
	}

	time.Sleep(time.Second)         // 主线程休眠一秒，如果不休眠，则主进程输出value时，可能部分协程还未进行完毕
	fmt.Println(c.Value("somekey")) // 输出该数值

	fmt.Println("=========================")

	var state = make(map[int]int) // int,int 的map
	var mutex = &sync.Mutex{}     // 锁对象
	var ops int64 = 0             // 原子计数，表示我们对状态进行了多少次操作
	var create int64 = 0

	for r := 0; r < 100; r++ {
		// 创建100个协程
		go func() {
			total := 0
			for {
				key := rand.Intn(5) // 随机0-5
				mutex.Lock()
				total += state[key] // 增加Total值
				//fmt.Println("Total:", total)
				mutex.Unlock()
				atomic.AddInt64(&ops, 1) // 增加原子计数
				runtime.Gosched()        // 释放切片时间
			}
		}()
	}
	for w := 0; w < 10; w++ {
		// 开启10个协程，随机生成map<int, int>表
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				//fmt.Println("key:", key, " val:", val)
				mutex.Unlock()
				atomic.AddInt64(&create, 1)

				runtime.Gosched() // 释放时间切片
			}
		}()
	}
	time.Sleep(time.Second) // 主线程休眠一秒，此时让协程去执行1秒

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
	createFinal := atomic.LoadInt64(&create)
	fmt.Println("create map times:", createFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
