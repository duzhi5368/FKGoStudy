package main

import (
	"fmt"
)

func Entering(s string) string {
	fmt.Println("Entering:", s)
	return s
}

func Leaving(s string) {
	fmt.Println("Leaving:", s)
}

func a() {
	defer Leaving(Entering("a"))
	fmt.Println("in a")
}

func b() {
	defer Leaving(Entering("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
}
