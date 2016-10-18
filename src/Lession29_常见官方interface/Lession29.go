package main

import (
	"fmt"
	"sort"
	"time"
)

type SPerson struct {
	Name string
	Age  int
}

// fmt库中的String() string 接口
func (p SPerson) String() string {
	//return "hi"
	return fmt.Sprintf("%v - %v years", p.Name, p.Age)
}

type SMyError struct {
	When time.Time
	What string
}

// error库中的Error() string 接口
func (e *SMyError) Error() string {
	return fmt.Sprintf("[%v]: %s", e.When, e.What)
}

// sort的三个自定义排序接口
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
	person1 := SPerson{"Demacia", 22}
	fmt.Println(person1) // 输出的是自定义格式

	fmt.Println("==========================")
	err := SMyError{time.Now(), "this is a testError"}
	fmt.Println(&err)

	fmt.Println("==========================")
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
