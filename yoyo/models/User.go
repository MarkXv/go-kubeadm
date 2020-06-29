package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


type User struct {
	Id int
	Name string
	Pwd string
}

func init()  {
	orm.RegisterDataBase("default","mysql","root:root@tcp(49.234.198.239:3306)/test?charset=utf8")
	//生成的表的名称
	orm.RegisterModel(new (User))
	//生成表，表别名/是否强制重新生成/是否可见
	orm.RunSyncdb("default",false,true)


}
