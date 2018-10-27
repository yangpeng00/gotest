package main

import (
	"beego-master"
	"beego-master/orm"
	"fmt"
)

type uer struct {
	id   int
	name string
	age  int
}

func init() {

	orm.RegisterModel(new(uer))
}

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/gotest?charset=utf8")
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	o := orm.NewOrm()
	o.Using("default")

	uer := new(uer)

	uer.name = "dsfs"
	uer.age = 23

	o.Insert(uer)
	fmt.Printf("%d 的年龄是 %v\n", uer.name, uer.age)
	beego.Run()
}
