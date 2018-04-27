package controllers

import (
	"github.com/astaxie/beego"
)
type MainController struct {
	beego.Controller
}
func (c *MainController) Chat() {
	c.TplName = "chatv2.tpl"
}