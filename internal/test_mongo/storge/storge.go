package storge

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

//@author xiaolan
//@lastUpdate 2019-08-05
//@comment
//数据库操作包
//目前仅支持mongodb数据库


var mMGO *mgo.Session
var mDBName string

// 初始化
func Init(_url string, _dbname string, _user string, _pass string) error {

	mDBName = _dbname

	// 连接数据库
	connect_err := _connect(_url, _user, _pass)
	if connect_err != nil {
		return errors.New(fmt.Sprintf("connect %v error: %v", _url, connect_err))
	}

	// 启动线程检测服务器状态
	go (func() {
		for {
			ping_err := mMGO.Ping()
			if ping_err != nil {
				fmt.Printf("数据库状态异常: %v 尝试重连!", ping_err)
				connect_err := _connect(_url, _user, _pass)
				if connect_err != nil {
					fmt.Printf("重新连接mongodb失败! 错误:%v!", connect_err)
				}
			}
			<-time.After(30 * time.Second)
		}
	})()

	return nil
}


//@author xiaolan
//@lastUpdate 2019-08-06
//@comment 连接mongodb数据库
//@param _url 连接mongodb连接
//@param _user 用户名
//@param _pass 密码
func _connect(_url string, _user string, _pass string) error {
	var connect_err error
	mMGO, connect_err = mgo.Dial(_url)
	if connect_err != nil {
		return connect_err
	}

	// 用户登陆
	if _user != "" {
		loginErr := mMGO.Login(&mgo.Credential{
			Username:    _user,
			Password:    _pass,
		})
		if loginErr != nil {
			return loginErr
		}
	}

	return nil
}



//@author xiaolan
//@lastUpdate 2019-08-05
//@comment 获取默认的MongoDB数据库指针
func GetDB() *mgo.Database {
	return mMGO.DB(mDBName)
}


//@author xiaolan
//@lastUpdate 2019-08-05
//@comment 获取指定数据库指针
func GetDBByName(_name string) *mgo.Database {
	return mMGO.DB(_name)
}



