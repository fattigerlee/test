package storge

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"testing"
)


type UserInfo struct {
	User string `bson:"user"`
	Pass string `bson:"pass"`
	Age uint32 `bson:"age"`
}

func TestInit(t *testing.T) {

	err := Init("mongodb://0.0.0.1:27017", "local_db", "", "")
	if err != nil {
		fmt.Printf(">> MongoDB: %v\n", err)
		return
	}

	//err = GetDB().C("sky_user").Insert(bson.M{
	//	"user": "asdasd002",
	//	"pass": "asdasd002",
	//	"age": uint32(skyCommon.RandInt(0, 100)),
	//})
	//if err != nil {
	//	fmt.Printf("添加数据失败! 错误: %v", err)
	//	return
	//}

	var __xx []UserInfo
	err = GetDB().C("sky_user").Find(bson.M{
		"age": bson.M{
			"$in": []uint64{88, 17, 95, 52},
		},
	}).All(&__xx)
	if err != nil {
		fmt.Printf(">> 获取用户数据失败! 错误: %v", err)
		return
	}

	for _, val:= range __xx {
		fmt.Printf(">> 获取用户: [%v] => %v (%v)\n", val.User, val.Pass, val.Age)
	}

}