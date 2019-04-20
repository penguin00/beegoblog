package models

import (
	"beegoblog/help"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

func GetAllCategories() (Categories []*Category, err error) {
	o := orm.NewOrm()
	Categories = make([]*Category, 0)
	_, err = o.QueryTable("category").All(&Categories)
	return Categories, err
}

func AddCategory(name string) error {

	if strings.TrimSpace(name) == "" {
		return nil
	}
	o := orm.NewOrm()
	cate := &Category{Title: name}

	err := o.QueryTable("category").Filter("title", name).One(cate)

	if err == nil {
		// error == nil  说明查询没有出错，说明找到了这条记录
		return err
	}

	cate.Created = help.TimeNow()
	_, err = o.Insert(cate)
	return err
}

func DeleteCategory(cid string) error {

	id, err := strconv.ParseInt(cid, 10, 64)

	if err != nil {
		return err
	}

	o := orm.NewOrm()
	cate := &Category{Id: id}
	_, err = o.Delete(cate)
	return err
}
