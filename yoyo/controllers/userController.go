package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"yoyo/models"
)

type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Get() {
	o := orm.NewOrm()

	user := models.User{
		//Id:   ,
		Name: "小明",
		Pwd:  "1111",
	}

	_,err := o.Insert(&user)
	if err != nil{
		beego.Info("error",err)
		return
	}
}
