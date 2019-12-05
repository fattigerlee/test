package util

import (
	"math/rand"
	"time"
)

func GenUniqueId() uint64 {
	// 循环500次
	uid := uint64(0)
	for i := 0; i < 100; i++ {
		uid = (uid + uint64(RandInt(0, 0xFFFFFFFFFFFF))) % 100000000000
	}
	return uid
}

func RandInt(_min int64, _max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return (rand.Int63() % (_max - _min)) + _min
}
