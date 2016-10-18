package main

import (
	"encoding/json"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

// 定义类
type Dog struct {
	Name  string
	Color string
}

func main() {

	// 链接MemCache服务器
	mc := memcache.New("127.0.0.1:11211")

	// 从MemCache中拉取数据
	fetchItem, err := mc.Get("dog")
	// 检查是否拉取得到数据
	if err != memcache.ErrCacheMiss {
		if err != nil {
			fmt.Println("Error fetching from memcache", err)
		} else {
			fmt.Println("Cache hit!")

			dog, err := DecodeData(fetchItem.Value)
			if err != nil {
				fmt.Println("Error decoding data from memcache", err)
			} else {
				fmt.Println("Dog name is: ", dog.Name)
			}
		}
	}

	// 创建一个对象
	spot := Dog{Name: "Spot", Color: "brown"}
	// 存储到MemCache中
	setItem := memcache.Item{Key: "dog", Value: EncodeData(spot), Expiration: 300}

	err = mc.Set(&setItem)
	if err != nil {
		fmt.Println("Error setting memcache item", err)
	}
}

// 转换为json进行读写
func DecodeData(raw []byte) (dog Dog, err error) {
	err = json.Unmarshal(raw, &dog)
	return dog, err
}

// 转换为json进行读写
func EncodeData(dog Dog) []byte {
	enc, err := json.Marshal(dog)
	if err != nil {
		fmt.Println("Error encoding Action to JSON", err)
	}
	return enc
}
