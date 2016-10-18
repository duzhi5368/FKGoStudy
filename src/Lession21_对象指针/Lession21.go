package main

import "fmt"

// 定义一个类
type Dog struct {
	Name  string
	Color string
}

func main() {

	// 创建类对象并赋属性
	Spot := Dog{Name: "Spot", Color: "brown"}

	// 获取对象指针
	SpotPointer := &Spot

	// 通过指针修改对象数据
	SpotPointer.Color = "black"

	fmt.Println(Spot.Color)

}
