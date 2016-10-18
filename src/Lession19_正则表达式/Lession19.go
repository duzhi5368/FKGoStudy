package main

import (
	"fmt"
	"regexp"
)

func basic_regexes() {
	// 创建一个查找数字的正则表达式
	pattern := "[0-9]+"
	str := "The 12 monkeys ate 48 bananas"

	// 编译正则表达式
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}
	if re.MatchString(str) {
		fmt.Println("Yes, matched a number")
	} else {
		fmt.Println("No, no match")
	}

	// 单次查询
	result_two := re.FindString(str)
	fmt.Println("Number matched:", result_two)

	// 全体查询
	results_three := re.FindAllString(str, 2)
	for i, v := range results_three {
		fmt.Printf("Match %d: %s\n", i, v)
	}

	// 正则替换
	results_four := re.ReplaceAllString(str, "xx")
	fmt.Println("Result:", results_four)
}

func case_insensitive() {
	// 和普通的正则表达式有些不一样，这里对大小写不敏感
	ptn := `(?i)^t.`
	str := "To be or not to be"

	re, err := regexp.Compile(ptn)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	result := re.FindString(str)
	fmt.Println("Result:", result)
}

func main() {
	basic_regexes()
	fmt.Println("------------")
	case_insensitive()
}
