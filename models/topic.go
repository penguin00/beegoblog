package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func GetAllTopics(isDesc bool) (topic []*Topic, err error) {
	topic = make([]*Topic, 0)
	o := orm.NewOrm()

	qs := o.QueryTable("topic")

	if isDesc {
		_, err = qs.OrderBy("-created").All(&topic)
	} else {
		_, err = qs.All(&topic)
	}

	return topic, err

}

func AddTopic(name, cate, content string) error {

	topic := &Topic{Title: name, Category: cate, Content: content}

	o := orm.NewOrm()
	err := o.QueryTable("topic").Filter("title", name).One(topic)
	if err != nil {
		return err
	}

	topic.Created = time.Now()
	topic.Updated = time.Now()
	_, err = o.Insert(topic)

	if err == nil {
		AllTopic := make([]*Topic, 0)
		_, err = o.QueryTable("topic").Filter("category", cate).All(&AllTopic)

		cate := new(Category)
		_ = o.QueryTable("category").Filter("title", cate).One(cate)
		cate.TopicCount = int64(len(AllTopic))
		_, err = o.Update(cate)
	}

	return err

}
