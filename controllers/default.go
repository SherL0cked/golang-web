package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) ReadBlogId() int64 {
	id := this.Ctx.Input.Param(":id")
	val, err := strconv.Atoi(id)
	if err != nil {
		this.Ctx.WriteString("get param id fail")
	}
	return int64(val)
}

func (this *MainController) AdminRedirect() {
	// time.Sleep(0.5 * time.Second)
	this.Redirect("/zfwadmin/catalog/", 302)
}
