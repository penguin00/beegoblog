package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	View            int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Lables          string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	View            int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

type Reply struct {
	Id      int64
	Tid     int64
	Name    string
	Content string `orm:"size(1000)"`
	Created time.Time
}

func init() {
	orm.RegisterModel(new(Category), new(Topic), new(Reply))
	_ = orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/beegoblog?charset=utf8")
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RunSyncdb("default", false, true)
	orm.Debug = true

}
