package controllers

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"strconv"

	"lottery/models"
	//"fmt"
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
	if name := c.GetString("name"); name != "" {
		admin.Name = name
	}

	if pwd := c.GetString("password"); pwd != "" {
		admin.Password = pwd
	}

	if role := c.GetString("role"); role != "" {
		admin.Role = role
	}
	
	if msg := models.Insert(&admin); msg != "" {
		c.Ctx.WriteString(msg)
	}else{
		c.Ctx.WriteString("success")
	}
}

// @router /update [get]
func (c *AdminController) Update() {
	var o = orm.NewOrm()
	o.Using("default")

	var admin models.User
	if name := c.GetString("name"); name != ""{
		admin.Name = name
	}

	if paw := c.GetString("password"); paw != ""{
		admin.Password = paw
	}
	
	if role := c.GetString("role"); role != ""{
		admin.Role = role
	}

	if msg := models.Update(&admin); msg != "" {
		c.Ctx.WriteString(msg)
	}else{
		c.Ctx.WriteString("success")
	}
}

// @router /delete [get]
func (c *AdminController) Delete() {
	var o = orm.NewOrm()
	o.Using("default")

	var admin models.User
	if name := c.GetString("name"); name != ""{
		admin.Name = name
	}

	if msg := models.Delete(&admin); msg != "" {
		c.Ctx.WriteString(msg)
	}else{
		c.Ctx.WriteString("success")
	}
}

// @router /query [get]
func (c *AdminController) Query() {
	var o = orm.NewOrm()
	o.Using("default")

	var admin models.User
	var name = c.GetString("name")

	if msg := admin.SelectPwdByName(name); msg != "" {
		c.Ctx.WriteString(msg)
	}else{
		c.Ctx.WriteString("success: " + admin.ToString())
	}
}
