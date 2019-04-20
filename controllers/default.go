package controllers

import (
	"beegoblog/models"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type BaseController struct {
	beego.Controller
	i18n.Locale
}
type HomeController struct {
	BaseController
}

func (c *BaseController) Prepare() {

	l := c.GetString("lang")

	if l == "zh-CN" {
		c.Lang = l
	} else {
		c.Lang = "en-US"
	}
}
func (c *HomeController) Get() {
	c.TplName = "home.html"
	c.Data["language"] = c.Tr("language")
	c.Data["isHome"] = true
	c.Data["isLogin"] = checkAccount(c.Ctx)
	cate := c.Input().Get("cate")
	lable := c.Input().Get("lable")
	c.Data["Topics"], _ = models.GetAllTopics(cate, lable, true)
	c.Data["Categories"], _ = models.GetAllCategories()
}
