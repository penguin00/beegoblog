package models

import (
	"beegoblog/help"
	"github.com/astaxie/beego/orm"
	"strconv"
)

func GetAllReply(tid string) ([]*Reply, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	reply := make([]*Reply, 0)
	_, err = o.QueryTable("reply").Filter("tid", tidNum).All(&reply)
	return reply, err
}

func AddReply(tid, name, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	reply := &Reply{Tid: tidNum, Name: name, Content: content, Created: help.TimeNow()}
	_, err = o.Insert(reply)
	err = UpdateTopicReplyCount(tidNum)
	return err
}

func DeleteReply(id string) error {
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	reply := &Reply{Id: idNum}
	_ = o.QueryTable("reply").Filter("id", idNum).One(reply)
	_, err = o.Delete(reply)
	err = UpdateTopicReplyCount(reply.Tid)
	return err
}

func UpdateTopicReplyCount(tid int64) error {

	AllReply := make([]*Reply, 0)
	o := orm.NewOrm()
	_, err := o.QueryTable("reply").Filter("tid", tid).OrderBy("-id").All(&AllReply)

	topic := new(Topic)
	_ = o.QueryTable("topic").Filter("id", tid).One(topic)
	topic.ReplyCount = int64(len(AllReply))
	topic.ReplyTime = AllReply[0].Created
	_, err = o.Update(topic)

	return err

}
