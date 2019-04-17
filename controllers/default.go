package controllers

import (
	"beegoblog/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	data := models.TestGet()
	c.Data["data"] = data

	c.TplName = "home.html"
}
