package main

import "fmt"

func GetSum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	GetSum(1, 2)
	GetSum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	GetSum(nums...)
}
