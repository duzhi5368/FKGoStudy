package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建数组，声明长度，无初始化
	var array1 = [5]string{}
	PrintStrArray(array1[:])

	// 创建数组，声明长度，局部顺序初始化
	var array2 = [5]string{"Fr", "Kni"}
	PrintStrArray(array2[:])

	// 创建数组，声明长度，指定元素初始化
	var array3 = [5]string{1: "Chris", 4: "Ron"}
	PrintStrArray(array3[:])

	// 创建数组，不声明长度
	var array4 = [...]string{"Hi", "Demacia"}
	PrintStrArray(array4[:])

	// 创建数组，不声明长度，指定元素初始化
	var array5 = [...]string{1: "Chris", 6: "Ron"}
	PrintStrArray(array5[:])

	// 创建数组切片
	var array6 = []string{"lala", "dede"}
	PrintStrArray(array6)

	// 声明新数组切片
	var array7 = [...]string{1: "Chris", 4: "dcde", 6: "Ron"}
	var slice1 = array7[2:5]
	PrintStrArray(slice1)

	var slice2 = array7[2 : len(array7)-1]
	PrintStrArray(slice2)

	// 证明切片是值，而非指针
	AddSlice(slice2)
	PrintStrArray(slice2)
	PrintStrArray(array7[:])

	fmt.Println("====================")

	// 注意，切片是地址，值。连copy都不是
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// 注意切片的默认值
	s := []int{2, 3, 5, 7, 11, 13}

	s1 := s[:2]
	fmt.Println(s1)
	s2 := s[1:]
	fmt.Println(s2)
	s3 := s[:]
	fmt.Println(s3)

	fmt.Println("====================")

	// 输出切片长度和占比
	intSlice := []int{2, 3, 5, 7, 11, 13}
	PrintSliceInfo(intSlice)

	intSlice = intSlice[:0]
	PrintSliceInfo(intSlice)

	intSlice = intSlice[:4]
	PrintSliceInfo(intSlice)

	intSlice = intSlice[2:]
	PrintSliceInfo(intSlice)

	fmt.Println("====================")

	// 空切片检查
	var emptySlice []int
	fmt.Println(emptySlice, len(emptySlice), cap(emptySlice))
	if emptySlice == nil {
		fmt.Println("nil!")
	}

	fmt.Println("====================")

	// make函数声明动态切片,append添加切片
	// make只能初始化创建标准类型map,slice,channel，返回的是对象引用
	// new和常规认知一样，返回的是自定义对象指针
	aSlice := make([]int, 5)
	PrintSliceInfoWithStr("a", aSlice)

	bSlice := make([]int, 0, 5)
	PrintSliceInfoWithStr("b", bSlice)

	cSlice := bSlice[:2]
	PrintSliceInfoWithStr("c", cSlice)

	dSlice := cSlice[2:5]
	PrintSliceInfoWithStr("d", dSlice)

	eSlice := append(dSlice, 2, 3, 4)
	PrintSliceInfoWithStr("e", eSlice)

	fmt.Println("====================")

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 切片的切片组
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

//输出字符串数组
func PrintStrArray(array []string) {
	for i, v := range array {
		fmt.Printf("index:%d  value:%s\t", i, v)
	}
	fmt.Println()
}

// 更变切换内容
func AddSlice(slice []string) {
	for i := 0; i < len(slice); i++ {
		slice[i] += "_1"
	}
}

// 输出切片信息
func PrintSliceInfo(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func PrintSliceInfoWithStr(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
