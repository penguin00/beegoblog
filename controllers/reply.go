package controllers

import (
	"beegoblog/models"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Add() {
	_ = models.AddReply(c.Input().Get("tid"), c.Input().Get("nickname"), c.Input().Get("content"))

	c.Redirect("/topic/view/"+c.Input().Get("tid"), 302)
	return
}

func (c *ReplyController) Delete() {
	if !checkAccount(c.Ctx) {
		return
	}

	_ = models.DeleteReply(c.Input().Get("rid"))

	c.Redirect("/topic/view/"+c.Input().Get("tid"), 302)
	return
}
