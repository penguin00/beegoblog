package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {

	if l.Input().Get("exit") == "true" {
		l.Ctx.SetCookie("uname", "", -1, "/")
		l.Ctx.SetCookie("pwd", "", -1, "/")
		l.Redirect("/", 302)
		return
	}
	l.TplName = "login.html"
	l.Data["isLogin"] = checkAccount(l.Ctx)

}

func (l *LoginController) Post() {

	uname := l.Input().Get("uname")
	pwd := l.Input().Get("pwd")
	autoLogin := l.Input().Get("autoLogin") == "on"

	if uname == beego.AppConfig.String("uname") &&
		pwd == beego.AppConfig.String("password") {

		maxAge := 0
		if autoLogin {
			maxAge = 3600
		}

		l.Ctx.SetCookie("uname", uname, maxAge, "/")
		l.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	l.Redirect("/login", 302)
	return

}

func checkAccount(ctx *context.Context) bool {
	uname, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	pwd, err := ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	return uname.Value == beego.AppConfig.String("uname") &&
		pwd.Value == beego.AppConfig.String("password")
}
