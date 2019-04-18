package controllers

import (
	"beegoblog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	c.TplName = "category.html"
	c.Data["isLogin"] = checkAccount(c.Ctx)
	c.Data["isCategory"] = true
	c.Data["categories"], _ = models.GetAllCategories()
}
func (c *CategoryController) Post() {

	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		_ = models.AddCategory(name)
	case "del":
		cid := c.Input().Get("id")
		_ = models.DeleteCategory(cid)

	}

	c.Redirect("/category", 302)
	return

}
