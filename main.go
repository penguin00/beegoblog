package main

import (
	_ "beegoblog/routers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"os"
)

func main() {
	//创建目录
	os.Mkdir("attachment", os.ModePerm)
	i18n.SetMessage("en-US", "conf/"+"locale_en-US.ini")
	i18n.SetMessage("zh-CN", "conf/"+"locale_zh-CN.ini")

	beego.Run()
}
