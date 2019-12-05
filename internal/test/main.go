package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"test/test/util"
)

type Test struct {
	a string
	b string
}

func (this *Test) Set(a string, b string) {
	this.a = a
	this.b = b
}

func main() {
	//test()

	//test1()

	//test2()

	//test3()

	//test4()

	//test5()

	//test6()

	test7()
}

func test() {
	a := 29
	b := 1 << 0
	permission := (a & b) == b

	fmt.Printf("a:%b\n", a)
	fmt.Println("结果:", permission)

	str := "\"abc"
	fmt.Println(str)

	fmt.Println(CheckString(str))

	fmt.Println("异或:", 3^1)
}

func test1() {
	a := "hello world"
	fmt.Println("替换:", strings.ReplaceAll(a, "o", "e"))
	fmt.Println("原始值:", a)
}

func test2() {
	var test Test
	testType := reflect.TypeOf(&test)
	fmt.Println(testType)

	for i := 0; i < testType.NumMethod(); i++ {
		method := testType.Method(i)
		fmt.Println("method:", method)
		fmt.Println("method type:", method.Type.NumIn())
		for i := 0; i < method.Type.NumIn(); i++ {
			fmt.Println("args:", method.Type.In(i))
		}
		method.Func.Call([]reflect.Value{reflect.ValueOf(&test), reflect.ValueOf("100"), reflect.ValueOf("200")})
	}

	//testValue := reflect.ValueOf(&test)
	//for i := 0; i < testValue.NumMethod(); i++ {
	//	method := testType.Method(i)
	//	fmt.Println("method:", method)
	//	fmt.Println("method type:", method.Type.NumIn())
	//	method.Func.Call([]reflect.Value{reflect.ValueOf(&test), reflect.ValueOf("100"), reflect.ValueOf("200")})
	//}

	fmt.Println(test)
}

func test3() {
	sha := GetMemberLockSha(5958297110, 88563430645, 2, "14789")
	fmt.Println("密码:", sha)

	sha = GetMemberLockSha(87147781280, 86314360364, 2, "1478")
	fmt.Println("密码:", sha)

	var num string
	list := strings.Split(num, ";")
	fmt.Println(list)
}

func GetMemberLockSha(_groupid uint64, _uid uint64, _lockType uint32, _locknum string) string {
	return Md5([]byte(fmt.Sprintf("%v-%v-%v-%v", _groupid, _uid, _lockType, _locknum)))
}

func Md5(str []byte) string {
	h := md5.New()
	h.Write(str) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func CheckString(_param string) string {
	match, err := regexp.Match(`[\;\'\"\<\>]`, []byte(_param))
	if err != nil || match {
		return fmt.Sprintf("检测字符串错误:%v", err)
	}
	return _param
}

func test4() {
	//url := "http://www.baidu.com"
	//url := "https://juejin.im/post/5aeb0e016fb9a07ab7740d90"
	url := "https://movie.douban.com"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("获取信息失败:", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("读取信息内容失败:", err)
		return
	}

	fmt.Println(string(data))
}

func test5() {
	type Test struct {
		a string
	}

	a := &Test{a: "hello world"}
	fmt.Println(a)

	b := a
	fmt.Println(b)

	b = nil
	fmt.Println(a)
	fmt.Println(b)
}

func test6() {
	dict := make(map[uint64]bool)

	for i := 0; i <= 100000000; i++ {
		uid := util.GenUniqueId()

		if _, ok := dict[uid]; ok {
			fmt.Println("uid已存在:", uid)
			break
		}
		dict[uid] = true
	}
	fmt.Println("无重复uid")
}

func test7() {
	s := "hello world"
	str := strings.Split(s, ",")
	fmt.Println(str)
	fmt.Println(len(str))
}
