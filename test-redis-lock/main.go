package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
)

// redis分布式锁实现
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("redis连接失败!!!")
	}
	fmt.Println("redis连接成功!!!")

}
