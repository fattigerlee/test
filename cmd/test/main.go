package main

import (
	"chess_game/pkg/common"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("uid:", GenUid(7))

	common.InitMongo()
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
