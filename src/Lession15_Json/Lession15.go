package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 声明一个类
type Person struct {
	Name, City string
}

// 创造类对象
var person Person

// 创建类对象数组
var people []Person

type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// JSON字符串
	json_str := "{ \"Name\": \"Marcus\", \"City\": \"San Jose\"}"
	// Unmarshal函数参数要求是 []byte ，所以我们将string转换过去，并填充给person对象
	if err := json.Unmarshal([]byte(json_str), &person); err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}
	fmt.Printf("Name: %v, City: %v\n", person.Name, person.City)

	//-----------------------
	// 批量加载
	//-----------------------
	// 加载一个Json文件
	file, err := ioutil.ReadFile("../../Res/Names.json")
	if err != nil {
		fmt.Println("Error reading file")
	}

	// 将Json中的数据填充给People对象数组
	if err := json.Unmarshal(file, &people); err != nil {
		fmt.Println("Error parsing JSON", err)
	}
	fmt.Println(people)

	// 使用Marshal函数向Json文件中写入一个对象数据
	json, err := json.Marshal(people)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	fmt.Println(string(json))

	fmt.Println("=====================")
	//-----------------------
	// JSON函数使用
	//-----------------------
	/*
		bolB, _ := json.Marshal(true)
		fmt.Println(string(bolB))
		intB, _ := json.Marshal(1)
		fmt.Println(string(intB))
		fltB, _ := json.Marshal(2.34)
		fmt.Println(string(fltB))
		strB, _ := json.Marshal("gopher")
		fmt.Println(string(strB))
		slcD := []string{"apple", "peach", "pear"}
		slcB, _ := json.Marshal(slcD)
		fmt.Println(string(slcB))
		mapD := map[string]int{"apple": 5, "lettuce": 7}
		mapB, _ := json.Marshal(mapD)
		fmt.Println(string(mapB))
		res1D := &Response1{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"}}
		res1B, _ := json.Marshal(res1D)
		fmt.Println(string(res1B))
		res2D := &Response2{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"}}
		res2B, _ := json.Marshal(res2D)
		fmt.Println(string(res2B))
		byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
		var dat map[string]interface{}
		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}
		fmt.Println(dat)
		num := dat["num"].(float64)
		fmt.Println(num)
		strs := dat["strs"].([]interface{})
		str1 := strs[0].(string)
		fmt.Println(str1)
		str := `{"page": 1, "fruits": ["apple", "peach"]}`
		res := Response2{}
		json.Unmarshal([]byte(str), &res)
		fmt.Println(res)
		fmt.Println(res.Fruits[0])
		enc := json.NewEncoder(os.Stdout)
		d := map[string]int{"apple": 5, "lettuce": 7}
		enc.Encode(d)
	*/
}
