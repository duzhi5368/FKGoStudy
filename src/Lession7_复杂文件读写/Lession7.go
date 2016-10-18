package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := "./../../Res/Input.txt"

	// 使用系统缓冲加载一个文件
	f, _ := os.Open(filepath)
	// 延迟释放
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// 输出文字
		line := scanner.Text()
		fmt.Println(line)
	}

	// 如果有错误则输出
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// 写文件
	f, err := os.Create("./../../Res/output.txt")
	if err != nil {
		log.Fatalln("Error creating file: ", err)
	}
	defer f.Close()

	// 写入字符串数组
	for _, str := range []string{"apple", "banana", "carrot"} {
		bytes, err := f.WriteString(str)
		if err != nil {
			log.Fatalln("Error writing string: ", err)
		}
		fmt.Printf("Wrote %d bytes to file\n", bytes)
	}

}
