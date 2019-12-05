package main

import (
	"fmt"
	"test/test_crypt/crypt"
	"time"
)

func main() {
	res := crypt.PKCS5Padding([]byte("a"))
	fmt.Println(len(res), res)

	data := []byte("我丢啊哈哈哈啊哈哈")
	key := []byte("1234561234512345")
	enc := crypt.Encode(data, key)
	fmt.Println("加密结果:", string(enc))

	time.Sleep(time.Second)
	dec := crypt.Decode(enc, key)
	fmt.Println("解密结果:", string(dec))
}
