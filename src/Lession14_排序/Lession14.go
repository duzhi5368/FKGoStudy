package main

import (
	"fmt"
	"sort"
)

// 这里是sort的三个接口，实现后可做自定义排序
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {

	// 源数据
	strEnglish := []string{"我好", "jkl", "ghi", "你好", "acd", "你不好", "jklq"}
	nums := []int{4, 2, 12, 5, 1, 3}
	fmt.Println("Raw ABC : ", strEnglish)
	fmt.Println("Raw Nums: ", nums)

	// string排序
	sort.Strings(strEnglish)
	fmt.Println("Sorted ABC:", strEnglish)

	// int排序
	sort.Ints(nums)
	fmt.Println("Sorted Nums:", nums)

	// 反向排序
	sort.Sort(sort.Reverse(sort.StringSlice(strEnglish)))
	fmt.Println("Reverse ABC:", strEnglish)

	//-------------------------------------------------
	// Map排序
	//-------------------------------------------------
	hash := map[string]int{
		"c": 3,
		"a": 6,
		"b": 2,
		"e": 5,
		"d": 4,
	}

	// 值得注意的是，此时的Map是完全乱序的，每次运行结果不同哟
	for k, v := range hash {
		fmt.Printf("%s => %v\t", k, v)
	}
	fmt.Println()

	// 根据Key排序
	keys := make([]string, 0, len(hash))
	for k := range hash {
		keys = append(keys, k)
	}
	// 这里仅仅对Key做了排序
	sort.Strings(keys)
	// 按照Key顺序进行输出，但注意，hash内部还是乱序的哟
	for i := range keys {
		fmt.Printf("%s => %v\t", keys[i], hash[keys[i]])
	}
	fmt.Println("")
	fmt.Println("================================")

	//-------------------------------------------------
	// 自定义函数
	//-------------------------------------------------

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
