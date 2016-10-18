package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将得到的值输入到管道c中
}

func main() {
	s := []int{1, 2, 0, -4, 4, 6, 8, 2, 4, 3}
	// make只能初始化创建标准类型map,slice,channel，返回的是对象引用
	// new和常规认知一样，返回的是自定义对象指针
	c := make(chan int)

	slice := s[:len(s)/2]
	for _, v := range slice {
		fmt.Printf("%v ", v)
	}
	fmt.Println("\n---------")
	go sum(slice, c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从管道中取得数据值
	fmt.Println(x, y, x+y)
}
