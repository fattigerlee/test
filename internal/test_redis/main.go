package main

import (
	"encoding/json"
	"fmt"
	"test/test_redis/redis"
)

type LevelType uint32

const (
	LevelTypeNone  LevelType = iota // 无类型
	LevelTypeSmall                  // 小
	LevelTypeBig                    // 大
)

type Student struct {
	Uid   uint64    `json:"uid"`
	Name  string    `json:"name"`
	Level LevelType `json:"level"`
}

func main() {
	Redis, err := redis.New("localhost:6379", "", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	stud := &Student{}
	stud.Uid = 200
	stud.Name = "jack"

	key := fmt.Sprintf("test:%v", stud.Uid)
	err = Redis.Set(key, stud, 1000)
	if err != nil {
		fmt.Println(err)
		return
	}

	res := Redis.Get(key)

	var stu Student
	_ = json.Unmarshal([]byte(res), &stu)

	fmt.Println(stu)
}
