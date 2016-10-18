// 文件读写操作测试
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	filename := "./../../Res/Input.txt"

	// 读取文件
	// 如果读取错误，将会返回Err; 若读取没有错误，Err将为nil
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		// 加载出错，则输出Log
		log.Fatalln("Error reading file", filename)
	}

	// 读取完毕的content内容是字节数组 []byte，所以我们先将其转换为string再输出
	fmt.Println(string(content))

	// 写入新文件
	outfile := "./../../Res/Output.txt"
	err = ioutil.WriteFile(outfile, content, 0644)
	if err != nil {
		log.Fatalln("Error writing file: ", err)
	}

}
