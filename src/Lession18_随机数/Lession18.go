package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MathFunc func(int, int) int

func main() {

	// 随机返回一个 [0,x) 的半开半闭区间内的int值
	// 注意，如果没有初始化种子，则这前三个值永远是 81,87,47
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println("--------------")

	// 添加了随机种子，则生成数开始随机
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println("--------------")

	// 其他一些随机方法，例如随机整形，随机Float
	// Perm 将0到N的数随机分布为一个数组
	fmt.Println("Random Int:", rand.Int())
	fmt.Println("Random Float:", rand.Float32())
	fmt.Println("Random Permutation of N ints:", rand.Perm(31))

	// 该NormFloat64随机数调用多次的话会生成一个曲线，不算标准随机
	for i := 0; i < 10; i++ {
		fmt.Println("Normalized:", rand.NormFloat64())
	}

	fmt.Println("------------------------")

	// 随机整数
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// 数据Float
	fmt.Println(rand.Float64())
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 设置种子
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 设置相同的种子之后，随机数是完全相同的
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()
}
