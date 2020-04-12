package routers

import (
	"systeMgr/app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.IndexController{})
}
