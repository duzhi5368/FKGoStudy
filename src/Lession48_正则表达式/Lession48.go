package main

import "bytes"
import "fmt"
import "regexp"

func main() {

	// 简单的检查字符串是否符合正则
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 其他复杂的正则表达式需要先进行编译
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))                   // 检查是否符合正则
	fmt.Println(r.FindString("peach punch"))              // 检查符合正则的部分
	fmt.Println(r.FindStringIndex("peach punch"))         // 检查符合正则部分的字符索引
	fmt.Println(r.FindStringSubmatch("peach punch"))      // 查找符合正则的子字符
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // 查找符合正则的子字符索引

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.Match([]byte("peach")))

	// 进行正则表达式的编译时，我们可以使用MustCompile替代Compile，它只有一个返回值
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
