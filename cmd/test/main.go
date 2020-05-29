package main

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"sort"
	"strings"
	"time"
)

func main() {
	//test01()
	test02()

	//getAddress()
	//
	//fmt.Println("时间:", time.Now().String())
	//
	//str := "abc"
	//newStr := strings.ReplaceAll(str, "b", "")
	//fmt.Println(newStr)
	//
	//fmt.Println("权限:", 1&1)
	//fmt.Println("大小:", 20<<20)
	//fmt.Println("时间:", time.Now().Unix())
	//fmt.Println("哈希:", GetChatId("abc"))
	//
	//a := 2
	//b := 2
	//c := a ^ b
	//fmt.Println("c:", c)
	//
	//fmt.Println("uid:", GenUid(7))
	//
	//uri := "https://192.168.1.11:8000/?p=abc"
	//ParseUri(uri)
	//
	//testSort()
}

// channel不关闭不会导致啥问题,但是如果这个对象在其他goroutine里被引用的话,这个goroutine就无法gc了,如果想要其他goroutine结束,你就close那个channel对象,其他goroutine会捕获一个关闭信号,然后退出就没事了
func test01() {
	var ch chan int
	go func() {
		ch = make(chan int, 10)
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//close(ch)
	}()

	time.Sleep(time.Second)
	go func() {
		for {
			data, ok := <-ch
			if !ok {
				fmt.Println("channel关闭了...")
				break
			}
			fmt.Println("读到数据了...", data)
		}
	}()

	time.Sleep(time.Second * 5)
}

func test02() {
	fmt.Println(31 & 2)
}

// 生成聊天唯一id
func GetChatId(sha string) string {
	now := time.Now().UnixNano() / 1e6 % 1e10
	return fmt.Sprintf("%s%d", sha, now)
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

func getAddress() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
