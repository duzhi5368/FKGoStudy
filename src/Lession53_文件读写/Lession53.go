package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//===============================
	// 读文件
	//===============================

	// 打开方式1
	dat, err := ioutil.ReadFile(getCurrentDirectory() + "/test")
	check(err)
	fmt.Print(string(dat))
	fmt.Println()
	// 打开方式2
	f, err := os.Open(getCurrentDirectory() + "/test")
	check(err)

	// 读取部分
	b1 := make([]byte, 6)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 跳过6字节
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 3)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// 跳过6字节
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 3)
	n3, err := io.ReadAtLeast(f, b3, 3)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 部分读取
	_, err = f.Seek(0, 0)
	check(err)
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(6)
	check(err)
	fmt.Printf("6 bytes: %s\n", string(b4))
	f.Close()

	//===============================
	// 写文件
	//===============================
	// 打开方式，并直接写入
	d1 := []byte("hello\ngo\n")
	err2 := ioutil.WriteFile(getCurrentDirectory()+"/test1", d1, 0644)
	check(err2)

	// 打开方式2
	f, err3 := os.Create(getCurrentDirectory() + "/test2")
	check(err3)

	// 强制延迟关闭
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n22, err4 := f.Write(d2)
	check(err4)
	fmt.Printf("wrote %d bytes\n", n22)

	n33, _ := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n33)
	f.Sync() // 写入到硬盘

	w := bufio.NewWriter(f)
	n44, _ := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n44)
	w.Flush() // 写入到硬盘
}
