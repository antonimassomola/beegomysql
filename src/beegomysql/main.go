package main

import (
	_ "beegomysql/routers"
	_ "beegomysql/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "beego:beego@tcp(172.26.0.20:3306)/beegomysql?charset=utf8")
}

func main() {
	beego.Run()
}

