package main

import "time"
import "fmt"

func main() {
	fmt.Println("Timer 1 创建")
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C // 执行等待，此处会因通道信息等待导致卡死
	fmt.Println("Timer 1 过期")

	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 过期") // 这里永远都不会被调用，因为主线程先已经退出了，协程被杀
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 停止")
	}
}
