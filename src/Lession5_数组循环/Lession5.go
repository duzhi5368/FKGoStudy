// 数组，循环
package main

import "fmt"

func main() {

	// 支持最基本的数组模式
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// 字符串数组
	alphas := []string{"abc", "def", "ghi"}

	// 标准循环
	// 注意：i++可以， ++i不可以。
	for i := 0; i < len(alphas); i += 1 {
		fmt.Printf("%d: %s \n", i, alphas[i])
	}

	// range关键字，返回两个关键字
	// i->索引编号，val->值
	for i, val := range alphas {
		fmt.Printf("%d: %s \t", i, val)
	}
	fmt.Println()

	// rang仅输出编号
	for i := range alphas {
		fmt.Println(i)
	}

	// 如果不关心索引编号，则用 _ 替代，这表示我们不需要保存该值
	for _, val := range alphas {
		fmt.Println(val)
	}

	// 我们也可以使用类似While的方式测试
	x := 0
	for x < 4 {
		fmt.Println(x)
		x++
	}

	// 将 go 字符串视为数组，range取值输出
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	// 无尽循环,去除测试后可调试
	//for {
	//	fmt.Print(".")
	//}
}
