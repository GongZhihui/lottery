package controllers

import(
	"github.com/astaxie/beego"
	//"strconv"
	"encoding/json"
	"lottery/models"
	//"fmt"
)

// AdminController .
type AdminController struct {
	beego.Controller
}

// Get .
// @Title Get
// @Description 初始化显示后台数据
// @Success 200 {object} models.LottoBackendView
// @router / [get]
func (c *AdminController) Get() {
	var back = models.MakeLottoBackend()
	back.QueryByID()
	//
	var view models.LottoBackendView
	view.Notice = "hello"
	view.LottoPrompt = "gzh"
	var item = make([]models.Item, 0)
	item = append(item, models.Item{"魔岩石", 0.5})
	item = append(item, models.Item{"金币", 0.25})
	item = append(item, models.Item{"无色", 0.01})
	item = append(item, models.Item{"牛奶", 0.02})
	item = append(item, models.Item{"龙人", 0.25})
	view.Items = item
	c.Data["json"] = view
	//c.Data["json"] = back.ToLottoBackendView()
	c.ServeJSON()
}

// Post .
// @Title Post
// @Description 提交后台数据
// @Param	body		body 	models.LottoBackendView	true  "body for LottoBackendView content"
// @Success 200 {bool} true
// @Failure 403 submit fail
// @router / [post]
func (c *AdminController) Post() {
	var view models.LottoBackendView
	json.Unmarshal(c.Ctx.Input.RequestBody, &view)
	c.Data["json"] = view
	c.ServeJSON()
}

// @router /insert [get]
func (c *AdminController) Insert() {
	 
}

// @router /update [get]
func (c *AdminController) Update() {
	
}

// @router /delete [get]
func (c *AdminController) Delete() {
	
}

// @router /query [get]
func (c *AdminController) Query() {
	
}
