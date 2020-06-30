package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"yoyo/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	o := orm.NewOrm()

	user := models.User{
		//Id:   ,
		Name: "小红",
		Pwd:  "1112",
	}

	_,err := o.Insert(&user)
	if err != nil{
		beego.Info("error",err)
		return
	}
	c.TplName = "user/userList.html"
}
