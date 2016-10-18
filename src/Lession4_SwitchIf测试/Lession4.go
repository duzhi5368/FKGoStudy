// switch,if测试
package main

import (
	"fmt"
	"time"
)

func main() {
	num := 3
	num2 := 2

	if num > 3 {
		fmt.Println("Many")
	}

	if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Many")
	}

	// 注意这里，switch的条件在case中，并且这里是顺序截断式的。
	// 这个是无条件switch
	// 另外，默认全部会break，无需额外编写break
	switch {
	case num == 1:
		fmt.Println("One")
	case num2 == 2:
		fmt.Println("Two")
	case num > 2:
		fmt.Println("Many")
	default:
		fmt.Println("Thrown over boat")
	}

	// 也支持常规switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}
}
