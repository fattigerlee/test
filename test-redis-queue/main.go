package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
	"sync"
	"time"
)

// redis消息队列实现
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("redis连接失败!!!")
	}
	fmt.Println("redis连接成功!!!")

	queueOne(rdb)

	queueMany(rdb)
}

// 消息队列一对一(list)
func queueOne(rdb *redis.Client) {
	var wg sync.WaitGroup

	wg.Add(2)

	// 生产者
	go func() {
		var count int
		for count < 3 {
			count++
			time.Sleep(time.Second)

			rdb.LPush("message", "this is a test message!!!")
		}
		rdb.LPush("message", "finish")
		wg.Done()
	}()

	// 消费者
	go func() {
		for {
			res, err := rdb.BRPop(0, "message").Result()
			if err != nil {
				fmt.Println("生产者list消息接收失败:", err)
				break
			}
			fmt.Println("收到生产者list消息:", res[1])

			if res[1] == "finish" {
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("测试完成!!!")
}

// 消息队列一对多(发布订阅)
func queueMany(rdb *redis.Client) {
	var wg sync.WaitGroup
	wg.Add(3)

	// 消费者1
	go func() {
		sub := rdb.Subscribe("message")
		for msg := range sub.Channel() {
			fmt.Println("收到生产者发布订阅消息1:", msg.Payload)
			if msg.Payload == "finish" {
				break
			}
		}
		wg.Done()
	}()

	// 消费者2
	go func() {
		sub := rdb.Subscribe("message")
		for msg := range sub.Channel() {
			fmt.Println("收到生产者发布订阅消息2:", msg.Payload)
			if msg.Payload == "finish" {
				break
			}
		}
		wg.Done()
	}()

	// 生产者
	go func() {
		var count int
		for count < 3 {
			count++
			time.Sleep(time.Second)
			rdb.Publish("message", "this is a test message!!!")
		}
		rdb.Publish("message", "finish")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("测试完成!!!")
}
