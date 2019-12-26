package test_crypt

import (
	"fmt"
	"test/internal/test_crypt/crypt"
	"time"
)

func main() {
	res := crypt.PKCS5Padding([]byte("a"))
	fmt.Println(len(res), res)

	data := []byte("我丢啊哈哈哈啊哈哈")
	key := []byte("0e67f635e4bbeb74b7609def8dab6e6e")
	enc := crypt.Encode(data, key)
	fmt.Println("加密结果:", string(enc))

	time.Sleep(time.Second)
	dec := crypt.Decode(enc, key)
	fmt.Println("解密结果:", string(dec))
}
