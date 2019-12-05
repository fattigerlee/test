package model_user

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"test/test_mongo/storge"
	"time"
)

const CollectionName = "user"

type LevelType uint32

const (
	LevelTypeNone  LevelType = iota // 无类型
	LevelTypeSmall                  // 小
	LevelTypeBig                    // 大
)

// 用户基本信息
type UserBase struct {
	Uid      uint64    `bson:"uid" json:"uid"`           // 唯一id
	Nickname string    `bson:"nickname" json:"nickname"` // 昵称
	Photo    string    `bson:"photo" json:"photo"`       // 头像
	Sex      uint32    `bson:"-" json:"sex"`             // 性别
	Level    LevelType `bson:"level"`                    // 等级
}

// 用户信息
type User struct {
	UserBase `bson:",inline"` // 用户基本信息
	Money    uint64           `bson:"-" json:"-"` // 金币
}

func (this *User) Create() {
	err := storge.GetDB().C(CollectionName).Insert(this)

	if err != nil {
		fmt.Println("创建用户信息失败:", err)
	}
}

func (this *User) Encode() interface{} {
	data, err := json.Marshal(this)
	if err != nil {
		return nil
	}
	return string(data)

	//return bson.M{
	//	"uid":      this.Uid,
	//	"nickname": this.Nickname,
	//	"photo":    this.Photo,
	//	"sex":      this.Sex,
	//	"money":    this.Money,
	//}
}

func (this *User) Encode1() interface{} {
	return bson.M{
		"uid":      this.Uid,
		"nickname": this.Nickname,
		"photo":    this.Photo,
		"sex":      this.Sex,
		"money":    this.Money,
	}
}

// 创建用户信息
func Create(uid uint64, nickname string, photo string) {
	user := NewUser()
	user.Uid = uid
	user.Nickname = nickname
	user.Photo = photo

	err := storge.GetDB().C("user").Insert(user)

	if err != nil {
		fmt.Println("创建用户信息失败:", err)
	}
}

// 删除用户信息
func Remove(uid uint64, nickname string) {
	err := storge.GetDB().C("user").Remove(bson.M{
		"uid":      uid,
		"nickname": nickname,
	})

	if err != nil {
		fmt.Println("删除用户信息失败:", err)
	}
}

// 批量删除用户信息
func RemoveMulti(uidList ...uint64) {
	_, err := storge.GetDB().C(CollectionName).RemoveAll(bson.M{
		"uid": bson.M{
			"$in": uidList,
		},
	})

	if err != nil {
		fmt.Println("批量删除用户信息失败:", err)
	}
}

// 删除过期时间
func RemoveExpire() {
	now := time.Now().Unix()
	expireTime := now

	_, err := storge.GetDB().C(CollectionName).RemoveAll(bson.M{
		"invite_time": bson.M{
			"$lte": expireTime,
		},
	})

	if err != nil {
		fmt.Println("删除过期时间失败:", err)
	}
}

func NewUser() *User {
	return &User{}
}
