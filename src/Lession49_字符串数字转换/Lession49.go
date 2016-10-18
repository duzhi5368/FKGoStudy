package main

import "strconv"
import "fmt"

func main() {
	// float64转换
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	// int转换
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	// int 16位转换
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	// uint转换
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// 简便接口
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	// 非法转换时的处理
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
