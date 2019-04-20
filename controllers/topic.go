package controllers

import (
	"beegoblog/models"
	"fmt"
	"github.com/astaxie/beego"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.TplName = "topic.html"
	c.Data["isLogin"] = checkAccount(c.Ctx)
	c.Data["isTopic"] = true
	c.Data["topics"], _ = models.GetAllTopics("", "", true)
}

func (c *TopicController) Post() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	tid := c.Input().Get("tid")
	title := c.Input().Get("title")
	cate := c.Input().Get("category")
	content := c.Input().Get("content")
	lables := c.Input().Get("lable")

	// 获取附件
	_, fh, _ := c.GetFile("attachment")

	var attachment string
	if fh != nil {
		// 保存附件
		attachment = fh.Filename
		_ = c.SaveToFile("attachment", path.Join("attachment", attachment))
	}
	//path.join 的作用是组合成 attachment/xx.png

	if len(tid) > 0 {
		_ = models.ModifyTopic(tid, title, cate, content, lables, attachment)
	} else {
		_ = models.AddTopic(title, cate, content, lables, attachment)
	}

	c.Redirect("/topic", 302)

}

func (c *TopicController) Add() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "topic_add.html"
	c.Data["isTopic"] = true
}

func (c *TopicController) Modify() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	tid := c.Input().Get("tid")
	c.Data["Topic"], _ = models.GetTopic(tid)
	c.TplName = "topic_modify.html"
	c.Data["isTopic"] = true
}

func (c *TopicController) View() {
	c.TplName = "topic_view.html"
	tid := c.Ctx.Input.Param("0")
	topices, _ := models.GetTopic(tid)
	lables := topices.Lables
	c.Data["Topic"] = topices
	c.Data["Lables"] = strings.Split(strings.Replace(strings.Replace(lables, "$", "", -1), "#", " ", -1), " ")
	c.Data["isTopic"] = true
	c.Data["isLogin"] = checkAccount(c.Ctx)
	c.Data["Replies"], _ = models.GetAllReply(tid)
}

func (c *TopicController) Delete() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	tid := c.Input().Get("tid")

	err := models.DeleteTopic(tid)
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect("/topic", 302)
	return
}
