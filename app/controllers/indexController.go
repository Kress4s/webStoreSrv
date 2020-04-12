package controllers

import (
	"github.com/astaxie/beego"
)

// URLMapping ...
func (c *IndexController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// IndexController ...
type IndexController struct {
	beego.Controller
}

// Get 首页
// @router /index [get]
func (c *IndexController) Get() {
	c.TplName = "index.html"
}
