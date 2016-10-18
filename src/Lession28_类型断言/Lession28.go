package main

import (
	"fmt"
)

func CheckType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Printf("unknown %T \n", v)
	}
}

func main() {
	var i interface{} = "hello"

	// 强制转换i为string类
	// 若失败，则会直接断言宕掉
	// 若成功，则正常赋值
	s := i.(string)
	fmt.Println(s)

	// 强制转换i为string类
	// 若失败，则OK为false，不会宕掉
	// 若成功，则OK为true，s正常赋值
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(int)
	fmt.Println(f, ok)

	//f := i.(int) // 这里会断言宕掉
	fmt.Println(f)

	fmt.Println("==============")

	// 类型检查
	CheckType(i)
}
