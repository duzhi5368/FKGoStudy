package main

import (
	"fmt"
	"math"
)

// 接口
// 实现多态的核心
type Abser interface {
	Abs() float64
}

// 实现MyFloat的Abs
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 实现Vertex的Abs
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser //  声明一个Abser格式的a

	f := MyFloat(-2.0)
	v := Vertex{3, 4}

	a = f // MyFloat类型的Abs实现了
	fmt.Println(a.Abs())

	a = &v // *Vertex类型的Abs实现了
	fmt.Println(a.Abs())

	//a = v // 这行会出错，因为Vertex类型的Ads没有被实现

}
