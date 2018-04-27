package routers

import (
	"github.com/astaxie/beego"
	"hdy/chat/controllers"
)
func init() {
	beego.Router("/", &controllers.MainController{}, "*:Chat")
}
