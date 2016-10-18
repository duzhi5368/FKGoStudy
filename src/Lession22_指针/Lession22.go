package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	i, j := 42, 2222

	p := &i
	fmt.Println(*p)
	*p = 21 // 修改的只是指针
	fmt.Println(i)

	// 打印指针
	fmt.Println(p)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	v := Vertex{1, 2}
	q := &v
	q.X = 1e9 // 类型被强制转换
	fmt.Println(v)
}
