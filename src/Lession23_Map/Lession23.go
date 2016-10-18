package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// make声明数组，数组Key为string, Value为Vertex结构
	m = make(map[string]Vertex)

	// 两种赋值方法
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Test Labs"] = Vertex{
		10.23, -9.23543,
	}

	var m2 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	// 数组的元素添加
	m2["Amazon"] = Vertex{1.2, -2.3}
	fmt.Println(m)
	fmt.Println(m2)

	fmt.Println("===================")

	m3 := make(map[string]int)

	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])
	// 数组的替换
	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])

	// 数组的删除
	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])

	v, ok := m3["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
