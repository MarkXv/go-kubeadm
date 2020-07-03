

// @APIVersion 1.0.1
// @Title beego Test API for 自定义
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// 配置路由映射,以上注解必须添加


package routers

import (
	"github.com/astaxie/beego"
	"yoyo/controllers"
)

func init() {
	ns := beego.NewNamespace("/",
		beego.NSNamespace("/",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
		beego.NSNamespace("/user/add",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
    //beego.Router("/", &controllers.MainController{})
    //beego.Router("/user/add",&controllers.UserController{})
}
