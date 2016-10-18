// 测试引用其他文件
package main

import (
	"FKCommonFunc"
	"fmt"
)

// 包级别的变量，同类型多变量声明，初始化设值
var value1, value2 int = 1, 2

// 短声明如下，不可在函数外
//value4 := 4

// 多值返回
func swap(x, y string) (string, string) {
	return y, x
}

// 裸返回，若返回值进行了命名，那么return时刻不声明返回值，且返回值在函数中可被使用
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 第三方文件的对象，直接和包名相同
func main() {
	fmt.Printf("Hi, func = %v\n", FKCommonFunc.Sqrt(2))

	// 注意，下面函数是无法执行的，原因是导出包里该函数的首字母是小写的
	// 只有首字母为大写的函数才会被导出
	// fmt.Printf("Hi, func = %v\n", FKCommonFunc.sqrt_inside(2))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	var value3 int = 17
	fmt.Println(split(value1))
	fmt.Println(split(value2))
	fmt.Println(split(value3))
}
