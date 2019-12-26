package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	fmt.Println("uid:", GenUid(7))

	uri := "https://192.168.1.11:8000/?p=abc"
	ParseUri(uri)

	testSort()
}

// 生成唯一id(默认六位数)
func GenUid(bitSize int) uint64 {
	if bitSize < 6 {
		return 0
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	min := uint64(math.Pow10(bitSize - 1))
	max := uint64(math.Pow10(bitSize))

	var uid uint64
	for {
		uid += r.Uint64()
		if uid%max > min {
			uid = uid % max
			break
		}
	}
	return uid
}

// 解析uri
func ParseUri(uri string) {
	params := strings.Split(uri, "?")
	if len(params) != 2 {
		fmt.Println("参数错误")
		return
	}

}

func testSort() {
	list := []int{1, 5, 2, 4, 5, 6, 3, 0}

	fmt.Println("before:", list)
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	fmt.Println("after:", list)
}
