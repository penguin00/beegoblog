package models

import (
	"beegoblog/help"
	"github.com/astaxie/beego/orm"
	"os"
	"path"
	"strconv"
	"strings"
)

func GetAllTopics(cate, lable string, isDesc bool) (topic []*Topic, err error) {
	topic = make([]*Topic, 0)
	o := orm.NewOrm()

	qs := o.QueryTable("topic")

	if len(cate) > 0 {
		qs = qs.Filter("category", cate)
	}
	if len(lable) > 0 {
		qs = qs.Filter("lables__contains", "$"+lable+"#")
	}

	if isDesc {
		_, err = qs.OrderBy("-created").All(&topic)
	} else {
		_, err = qs.All(&topic)
	}

	return topic, err

}

func GetTopic(tid string) (*Topic, error) {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	err = o.QueryTable("topic").Filter("id", id).One(topic)
	if err != nil {
		return nil, err
	}
	topic.View++
	_, err = o.Update(topic)

	return topic, err

}

func DeleteTopic(tid string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: id}
	err = o.QueryTable("topic").Filter("id", id).One(topic)

	_, err = o.Delete(topic)
	err = UpdateCategoryTopicCount(topic.Category)
	return err
}

func AddTopic(name, cate, content, lables, attachment string) error {

	topic := &Topic{
		Title:      name,
		Category:   cate,
		Content:    content,
		Lables:     strLabels(lables),
		Attachment: attachment,
		Created:    help.TimeNow(),
		Updated:    help.TimeNow(),
	}

	o := orm.NewOrm()
	err := o.QueryTable("topic").Filter("title", name).One(topic)
	if err == nil {
		return nil
	}

	_, err = o.Insert(topic)

	if err == nil {
		err = UpdateCategoryTopicCount(cate)
	}
	return err
}
func ModifyTopic(tid, name, cate, content, lables, attachment string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	Topic := &Topic{Id: id}
	o := orm.NewOrm()
	//err = o.QueryTable("topic").Filter("id", id).One(Topic)
	if o.Read(Topic) == nil {
		OldCate := Topic.Category
		OldAttachment := Topic.Attachment
		Topic.Title = name
		Topic.Category = cate
		Topic.Content = content
		Topic.Updated = help.TimeNow()
		Topic.Lables = strLabels(lables)

		if len(attachment) > 0 {
			Topic.Attachment = attachment
			_ = os.Remove(path.Join("attachment", OldAttachment))
		}
		_, err = o.Update(Topic)
		if cate != OldCate && err == nil {
			err = UpdateCategoryTopicCount(cate)
			err = UpdateCategoryTopicCount(OldCate)
		}
		return err
	}
	return nil

}

func UpdateCategoryTopicCount(cate string) error {
	if len(cate) > 0 {
		AllTopic := make([]*Topic, 0)
		o := orm.NewOrm()
		_, err := o.QueryTable("topic").Filter("category", cate).All(&AllTopic)

		category := new(Category)
		_ = o.QueryTable("category").Filter("title", cate).One(category)
		category.TopicCount = int64(len(AllTopic))
		_, err = o.Update(category)

		return err
	}
	return nil
}
func strLabels(labels string) string {
	if len(labels) > 0 {
		return "$" + strings.Join(strings.Split(labels, " "), "#$") + "#"
	}
	return ""

}
