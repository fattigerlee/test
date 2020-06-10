package main

import (
	rand2 "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

func main() {
	//test01()
	//test02()

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

	//test10()
	//test11()
	//test12()
	//test13()
	//test14()
	//test15()
	//test16()
	//test17()
	//test18()
	test20()
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

type S struct{}

func (s S) F() {}

type IF interface {
	F()
}

func InitType() S {
	var s S
	return s
}

func InitPointer() *S {
	var s *S
	return s
}
func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

func InitIfaceType() IF {
	var s S
	return s
}

func InitIfacePointer() IF {
	var s *S
	return s
}

func test10() {
	s := []int{1, 2, 3}
	ss := s[1:]
	fmt.Printf("ss大小:%d ss:容量:%d\n", len(ss), cap(ss))
	ss = append(ss, 4)
	fmt.Printf("ss大小:%d ss:容量:%d\n", len(ss), cap(ss))

	for _, v := range ss {
		v += 10
	}

	for i := range ss {
		ss[i] += 10
	}

	fmt.Println(s)
}

func test11() {
	//println(InitType() == nil)
	println(InitPointer() == nil)
	a := InitEfaceType()
	fmt.Println(a)
	println(InitEfaceType() == nil)
	println(InitEfacePointer() == nil)
	println(InitIfaceType() == nil)
	println(InitIfacePointer() == nil)
}

type SS struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func test12() {
	s := SS{}
	p := &s

	f(s) //A
	//g(s) //B
	f(p) //C
	//g(p) //D
}

func test13() {
	m := make(map[int]*int)

	for i := 0; i < 3; i++ {
		m[i] = &i //A
	}

	for _, v := range m {
		print(*v)
	}
}

func test14() {
	f, err := os.Open("file")
	if err != nil {
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	println(string(b))
}

func f1() {
	defer println("f1-begin")
	f2()
	defer println("f1-end")
}

func f2() {
	defer println("f2-begin")
	f3()
	defer println("f2-end")
}

func f3() {
	defer println("f3-begin")
	panic(0)
	defer println("f3-end")
}

func test15() {
	f1()
}

func test16() {
	const N = 10

	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()
	println(len(m))
}

type S1 struct{}

func (s1 S1) f() {
	fmt.Println("S1.f()")
}
func (s1 S1) g() {
	fmt.Println("S1.g()")
}

type S2 struct {
	S1
}

func (s2 S2) f() {
	fmt.Println("S2.f()")
}

type I interface {
	f()
}

func printType(i I) {
	if s1, ok := i.(S1); ok {
		s1.f()
		s1.g()
	}
	if s2, ok := i.(S2); ok {
		s2.f()
		s2.g()
	}
}

func test17() {
	printType(S1{})
	printType(S2{})
}

// 下载文件
func test18() {
	resp, err := http.PostForm("http://192.168.1.11:18888/download", url.Values{"skey": {"808a26bcb5a8b716a8abb20167300e15"}})
	//resp, err := http.PostForm("http://47.75.51.226:18888/download", url.Values{"skey": {"808a26bcb5a8b716a8abb20167300e15"}})
	//resp, err := http.PostForm("http://47.75.51.226:18888/download", url.Values{"skey": {"226d0f8c53060fa645074036ba046e8f"}})
	if err != nil {
		log.Fatal("下载文件失败,错误:", err)
	}
	defer resp.Body.Close()

	now := time.Now()
	var totalBuf []byte
	var i int
	for {
		i++
		buf := make([]byte, 32*1024)
		n, err := resp.Body.Read(buf)
		if n > 0 {
			totalBuf = append(totalBuf, buf[0:n]...)
		}

		if err != nil {
			break
		}
		//fmt.Printf("下载第%d次", i)
	}
	fmt.Printf("耗时:%.2fs", time.Since(now).Seconds())
	//fmt.Println(string(totalBuf))
}

func test19() {
	b := make([]byte, 16)
	_, err := rand2.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value:", base64.StdEncoding.EncodeToString(b))
}

func test20() {
	a := int(time.Hour / time.Millisecond)
	fmt.Println("a:", a)

	for i := 0; i < 10; i++ {
		test19()
	}
}

func test21() {
	json.Marshal(nil)
}
