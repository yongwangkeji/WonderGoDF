package database

import (
	"log"

	"gopkg.in/mgo.v2"
)

var MongodbSession *mgo.Session
var Err error

// 读取配置连接 mongodb 数据库
func ConnectMongoDB(addr string) {
	session, err := mgo.Dial(addr)
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	log.Println("connect to mongodb success")

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	MongodbSession = session
}

// 设置 Mongodb 数据库和表
func SetDBTable(DB string, Table string) *mgo.Collection {
	c := MongodbSession.DB(DB).C(Table)
	return c
}
