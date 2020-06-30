
package routers

import (
	"github.com/astaxie/beego"
	"yoyo/controllers"
)

func init() {
	//ns := beego.NewNamespace("/v1",
	//	beego.NSNamespace("/object",
	//		beego.NSInclude(
	//			&controllers.UserController{},
	//		),
	//	),
	//	beego.NSNamespace("/user",
	//		beego.NSInclude(
	//			&controllers.UserController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns)
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user/add",&controllers.UserController{})
}
