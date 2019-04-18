package controllers

import (
	"beegoblog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.TplName = "topic.html"
	c.Data["isLogin"] = checkAccount(c.Ctx)
	c.Data["isTopic"] = true
	c.Data["topics"], _ = models.GetAllTopics(true)
}
