package routers

import (
	"beegoblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	beego.AutoRouter(&controllers.ReplyController{})

	//controller 处理
	beego.Router("/attachment/:all", &controllers.AttachController{})
	//静态文件处理
	//beego.SetStaticPath("/attachment", "attachment")
}
