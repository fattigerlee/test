package model_setting

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"test/test_mongo/storge"
)

const COLLECTION_NAME = "sky_system_setting"

type ModelSystemSetting struct {
	Uid           uint64 `bson:"uid" json:"-"`
	NewMsgNotice  uint32 `bson:"newmsg_notice" json:"newmsg_notice"`
	RtcNotice     uint32 `bson:"rtc_notice" json:"rtc_notice"`
	RtcVoice      uint32 `bson:"rtc_voice" json:"rtc_voice"`
	NoticeDetail  uint32 `bson:"notice_detail" json:"notice_detail"`
	NotDisturb    uint32 `bson:"not_disturb" json:"not_disturb"`
	Voice         uint32 `bson:"voice" json:"voice"`
	Shock         uint32 `bson:"shock" json:"shock"`
	RecommentBook uint32 `bson:"recomment_book" json:"recomment_book"`
	CircleOpen    uint32 `bson:"circle_open" json:"circle_open"`
	Language      uint32 `bson:"language" json:"language"`
	FontSize      uint32 `bson:"fontsize" json:"fontsize"`
	ChatBG        string `bson:"chat_bg" json:"chat_bg"`
	AssetAutoLoad uint32 `bson:"asset_autoload" json:"asset_autoload"`
	EarpieceMode  uint32 `bson:"earpiece_mode" json:"earpiece_mode"`
	Skin          string `bson:"skin" json:"skin"`
}

//@author xiaolan
//@lastUpdate 2019-09-20
//@comment 设置系统配置到数据库中
//@param _setting 要设置的配置结构体
func SetSystemSetting(_setting ModelSystemSetting) {
	data, err := bson.Marshal(&_setting)
	if err != nil {
		fmt.Println(err)
		return
	}

	m := bson.M{}
	err = bson.Unmarshal(data, m)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = storge.GetDB().C(COLLECTION_NAME).Upsert(bson.M{
		"uid": _setting.Uid,
	}, bson.M{
		"$set": m,
	})

	if err != nil {
		fmt.Println(err)
	}
}

// 修改
func SetNewMsgNotice(uid uint64, newMsgNotice uint32) {
	_, err := storge.GetDB().C(COLLECTION_NAME).Upsert(bson.M{
		"uid": uid,
	}, bson.M{
		"$set": bson.M{
			"newmsg_notice": newMsgNotice,
		},
	})

	if err != nil {
		fmt.Println("SetNewMsgNotice:", err)
	}
}

// 获取
func GetSystemSetting(uid uint64) *ModelSystemSetting {
	var this ModelSystemSetting

	err := storge.GetDB().C(COLLECTION_NAME).
		Find(bson.M{"uid": uid}).
		One(&this)

	if err != nil {
		fmt.Println("GetSystemSetting:", err)
		return nil
	}
	return &this
}
