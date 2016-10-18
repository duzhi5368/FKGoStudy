// string的功能测试
package main

// 注意：如果你引用了一个库，但实际上却没有使用它，编译器会报错。
import (
	"fmt"     // 标准输出库
	"strings" // 字符串操作
)

// Main
func main() {
	// 小写转换
	str := "HI, I'M UPPER CASE"
	lower := strings.ToLower(str)
	fmt.Println(lower)

	// 字符串包含测试
	if strings.Contains(lower, "case") {
		fmt.Println("Yes, exists!")
	}

	// 字符串截取
	fmt.Println("Characters 3-9: " + str[3:9])
	fmt.Println("First Five: " + str[:5])

	// 字符串拆离
	sentence := "I'm a sentence made up of words"
	words := strings.Split(sentence, " ")

	// 格式化 %v 是指输出数组
	fmt.Printf("%v \n", words)

	// 如果使用空格进行字符串拆分的话，使用Fields函数会更好
	// 因为这样不仅仅可以使用空格拆分，还可以支持所有的空白字符拆分
	fields := strings.Fields(sentence)
	fmt.Printf("%v \n", fields)

	fmt.Println("==================")

	fmt.Println("Contains:  ", strings.Contains("test", "es"))
	fmt.Println("Count:     ", strings.Count("test", "t"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))
	fmt.Println("Index:     ", strings.Index("test", "e"))
	fmt.Println("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", strings.Repeat("a", 5))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", strings.ToLower("TEST"))
	fmt.Println("ToUpper:   ", strings.ToUpper("test"))
	fmt.Println("Len: ", len("hello"))
	fmt.Println("Char:", "hello"[1])
	fmt.Println()
}
