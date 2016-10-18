package main

import (
	"flag"
	"fmt"
	"os"
)

var str string
var num int
var help bool

func main() {

	// 获取控制台参数列表
	num_args := len(os.Args)

	// 检查参数个数是否合理
	if num_args < 2 {
		fmt.Println(">> No args passed in")
	}

	// flag包，包括了命令行参数的解析功能
	// 这个例子中我们创建了全局变量来存储这些参数
	flag.StringVar(&str, "str", "default value", "text description")
	flag.IntVar(&num, "num", 5, "text description")
	flag.BoolVar(&help, "help", false, "Display Help")
	flag.Parse()

	// 检查Help是否为true可显示
	if help {
		fmt.Println(">> Display help screen")
		os.Exit(1)
	}

	fmt.Println(">> String:", str)
	fmt.Println(">> Number:", num)
	// 输出命令行的最后一个参数
	fmt.Println(">> Last Item:", os.Args[num_args-1])

	// 在调用flag.Parse()之后，flag.Args就保存了命令行参数列表
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println(">> Flag Arg:", args[0])
	}

	// 另外一套方式获取参数处理
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := ""
	if len(os.Args) >= 3 {
		arg = os.Args[3]
	} else {
		arg = "参数不足3位"
	}
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
