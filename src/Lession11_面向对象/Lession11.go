package main

import (
	"fmt"
)

// 定义一个类，狗类
type Dog struct {
	Name  string
	Color string
}

// 使用对象的函数
func (d Dog) Call() {
	fmt.Printf("Come here %s dog, %s \n", d.Color, d.Name)
}

func main() {
	// 创建对象实例，并设置其属性
	SpotDogInstance := Dog{Name: "Spot", Color: "brown"}
	// 对象调用
	SpotDogInstance.Call()
}
