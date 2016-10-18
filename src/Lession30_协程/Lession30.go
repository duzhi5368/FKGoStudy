package main

import (
	"fmt"
	"time"
)

func printWord(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go printWord("Goroutines")
	printWord("MainThread")
}
