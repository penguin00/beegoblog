package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

func GetAllCategories() (Categories []*Category, err error) {
	return nil, nil
}

func AddCategory(name string) error {

	o := orm.NewOrm()
	cate := &Category{Title: name}

	err := o.QueryTable("category").Filter("title", name).One(cate)

	if err != nil {
		return err
	}

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
