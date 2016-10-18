package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {
	if arg == 42 {
		// 构造一个新错误并返回
		return -1, errors.New("can't work with 42")
	}

	// 无错误则返回nil
	return arg + 3, nil
}

// 自定义错误
type argError struct {
	arg  int
	prob string
}

// 重载error 的 interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// 自定义错误，则这样返回错误
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// 测试标准错误
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	// 测试自定义错误
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 注意下文的if使用
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
