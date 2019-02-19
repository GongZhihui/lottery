package controllers

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"

	"lottery/utils"
	"lottery/models"
	"fmt"
)

// AdminController .
type AdminController struct {
	beego.Controller
}


// @router /insert [get]
func (c *AdminController) Insert() {
	var o = orm.NewOrm()
	o.Using("default")

	var admin models.User
	admin.ID = utils.GenerateRandomString(8)
	admin.Name = "li1"
	admin.Password = "123"
	admin.Role = "admin"
	i , err := o.Insert(&admin)
	var msg string
	if err != nil{
		msg = err.Error()		
	}else{
		msg = ""
	}
	c.Ctx.WriteString(strconv.Itoa(int(i)) + msg)
}

// @router /update [get]
func (c *AdminController) Update() {
	var o = orm.NewOrm()
	o.Using("default")

	var admin models.User
	if id := c.GetString("id"); id != ""{
		admin.ID = id
	}
	if name := c.GetString("name"); name != ""{
		admin.Name = name
	}

	if paw := c.GetString("password"); paw != ""{
		admin.Password = paw
	}
	
	if role := c.GetString("role"); role != ""{
		admin.Role = role
	}

	fmt.Println(admin.ToString())

	i , err := o.Update(&admin)

	var msg string
	if err != nil{
		msg = err.Error()		
	}else{
		msg = ""
	}
	c.Ctx.WriteString(strconv.Itoa(int(i)) + msg)
}

