package controllers

import (
	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
}

func (c *baseController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "_index.html"
}
