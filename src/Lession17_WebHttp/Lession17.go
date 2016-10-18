package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Http负责处理的请求
	http.HandleFunc("/hello", helloRequest)
	http.HandleFunc("/", getRequest)

	// 启动Web服务器
	log.Println("Listening on http://localhost:9999/")
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// 处理 /hello 请求
func helloRequest(w http.ResponseWriter, r *http.Request) {

	// 注意，这里用的是Fprint，而不是Printf
	fmt.Fprint(w, "Hallo Welt")
	return
}

// this function simply serves up the file requested, or
// an index list if a directory is requested
func getRequest(w http.ResponseWriter, r *http.Request) {
	file_requested := "./" + r.URL.Path
	http.ServeFile(w, r, file_requested)
	return
}
