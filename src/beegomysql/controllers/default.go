package controllers

import (
	"github.com/astaxie/beego"
	"beegomysql/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Test"] = "Test string"
	c.TplName = "index.tpl"
}


func (c *MainController) UserRegister() {

	username := c.Ctx.Input.Param(":username")

	err := models.AddUser(username)

	c.Data["Username"] = username
	c.Data["Result"] = "Created"

	if(err != nil){
		c.Data["Result"] = "Created"
	}

	c.TplName = "register.tpl"

}