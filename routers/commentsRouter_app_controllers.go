package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["systeMgr/app/controllers:IndexController"] = append(beego.GlobalControllerRouter["systeMgr/app/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
