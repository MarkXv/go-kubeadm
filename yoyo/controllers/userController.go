package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"yoyo/models"
)

type UserController struct {
	beego.Controller
}

type JSON struct {
	Code string
	Mag string
	User string `json:"user_info"`

}

// swagger注解配置
// @Title Get
// @Description get user
// @Success 200 {object} models
// @Failure 403 error
// @router / [Get]
func (c *UserController) Get() {
	//o := orm.NewOrm()

	//user := models.User{
	//	//Id:   ,
	//	Name: "小红",
	//	Pwd:  "1112",
	//}

	//_,err := o.Insert(&user)
	//if err != nil{
	//	beego.Info("error",err)
	//	return
	//}
	//
	//res,err := o.QueryTable("user")
	//c.TplName = "user/userList.html"

	//b,err := json.Marshal(user)
	//if err != nil {
	//	fmt.Println("序列化异常")
	//}

	//查询
	//o := orm.NewOrm()
	//
	//user := models.User{}
	//user.Id = 1
	//
	//err := o.Read(&user)
	//
	//if err != nil{
	//	beego.Info("field")
	//}
	//beego.Info(user,"----sucess")


	//更新

	o := orm.NewOrm()

	user := models.User{}

	user.Id = 1

	err := o.Read(&user)

	if err == nil {
		user.Name = "test"
		o.Update(&user)
	}
	o.Read(&user)
	beego.Info(user)



	//c.Data["json"] = data
	c.ServeJSON()
}
