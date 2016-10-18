package main

import (
	"fmt"
)

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	// Defer其实是一个压栈行为
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("End")
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (result int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (result int) {
	t := 5
	defer func() {
		t = t + 5
		result += 1
	}()
	return t
}

func f4() (result int) {
	defer func(result int) {
		result = result + 5
	}(result)
	return 1
}
