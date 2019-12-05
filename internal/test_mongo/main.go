package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"test/test_mongo/model/model_user"
	"test/test_mongo/storge"
)

type SkyResp struct {
	Resp string      `json:"resp"` // 回传的响应
	Code uint32      `json:"code"` // 0=成功, >0为失败
	Msg  string      `json:"msg"`  // 失败提示内容
	Data interface{} `json:"data"` // 一些扩展数据
}

func main() {
	//test()

	//test1()

	//test2()

	test3()
}

// 测试返回
func test() {
	user := model_user.NewUser()
	user.Uid = 100
	user.Nickname = "哈哈哈"
	user.Sex = 1
	user.Money = 10000

	resp, _ := json.Marshal(SkyResp{
		Resp: "",
		Code: 0,
		Msg:  "",
		Data: user.Encode(),
	})

	fmt.Println("返回信息:", string(resp))

	resp1, _ := json.Marshal(SkyResp{
		Resp: "",
		Code: 0,
		Msg:  "",
		Data: user.Encode1(),
	})

	fmt.Println("返回信息:", string(resp1))

	user2 := model_user.NewUser()
	user2.Uid = 2
	user2.Nickname = "呵呵呵"

	var oldFriendResp []*model_user.User
	oldFriendResp = append(oldFriendResp, user, user2)

	resp2, _ := json.Marshal(SkyResp{
		Resp: "",
		Code: 0,
		Msg:  "",
		Data: oldFriendResp,
	})
	fmt.Println("返回信息:", string(resp2))

	// resp3
	resp3, _ := json.Marshal(SkyResp{
		Resp: "",
		Code: 0,
		Msg:  "",
		Data: user,
	})

	fmt.Println("返回信息:", string(resp3))
}

// 创建索引
func test1() {
	err := storge.Init("localhost:27017", "test", "", "")
	if err != nil {
		log.Fatal(err)
	}

	var index mgo.Index
	index.Key = []string{"uid", "nickname"}
	index.Unique = true
	index.Background = true

	err = storge.GetDB().C(model_user.CollectionName).EnsureIndex(index)
	if err != nil {
		fmt.Println("创建索引失败:", err)
	}

	//model_user.Create(100, "abc", "111")
	model_user.Create(200, "abc", "111")
}

func test2() {
	err := storge.Init("localhost:27017", "test", "", "")
	if err != nil {
		log.Fatal(err)
	}

	user := &model_user.User{}
	user.Uid = 12345
	user.Level = model_user.LevelTypeBig

	user.Create()
}

func test3() {
	err := storge.Init("localhost:27017", "test", "", "")
	if err != nil {
		log.Fatal(err)
	}

	model_user.RemoveMulti(2)
	//model_user.RemoveExpire()

}
