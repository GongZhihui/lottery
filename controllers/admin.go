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
	var view = back.ToLottoBackendView()
	// view.Notice = "hello"
	// view.LottoPrompt = "gzh"
	// var item = make([]models.Item, 0)
	// item = append(item, models.Item{"魔岩石", 0.5})
	// item = append(item, models.Item{"金币", 0.25})
	// item = append(item, models.Item{"无色", 0.01})
	// item = append(item, models.Item{"牛奶", 0.02})
	// item = append(item, models.Item{"龙人", 0.25})
	// view.Items = item
	c.Data["json"] = view
	//c.Data["json"] = back.ToLottoBackendView()
	c.ServeJSON()
}

// Notice .
// @Title Post
// @Description 提交公告
// @Param	body		body 	models.LottoBackendView	true  "body for LottoBackendView.Notice content"
// @Success 200 {bool} true
// @Failure 403 submit fail
// @router /notice [post]
func (c *AdminController) Notice() {
	var view models.LottoBackendView
	json.Unmarshal(c.Ctx.Input.RequestBody, &view)

	var backend = models.MakeLottoBackend()
	backend.Notice = view.Notice

	models.Update(&backend, "Notice")
	c.Data["json"] = true
	c.ServeJSON()
}

// Prompt .
// @Title Post
// @Description 提交提示
// @Param	body		body 	models.LottoBackendView	true  "body for LottoBackendView.LottoPrompt content"
// @Success 200 {bool} true
// @Failure 403 submit fail
// @router /prompt [post]
func (c *AdminController) Prompt() {
	var view models.LottoBackendView
	json.Unmarshal(c.Ctx.Input.RequestBody, &view)

	var backend = models.MakeLottoBackend()
	backend.LottoPrompt = view.LottoPrompt

	models.Update(&backend, "LottoPrompt")
	c.Data["json"] = true
	c.ServeJSON()
}

// LottoItem .
// @Title Post
// @Description 提交奖品项
// @Param	body		body 	models.LottoBackendView	true  "body for LottoBackendView.Items content"
// @Success 200 {bool} true
// @Failure 403 submit fail
// @router /lotto_item [post]
func (c *AdminController) LottoItem() {
	var view models.LottoBackendView
	json.Unmarshal(c.Ctx.Input.RequestBody, &view)

	var backend = models.MakeLottoBackend()
	backend.Items = view.SerializeItem()

	models.Update(&backend, "Items")
	c.Data["json"] = true
	c.ServeJSON()
}

// @router /delete [get]
func (c *AdminController) Delete() {
	
}

// @router /query [get]
func (c *AdminController) Query() {
	
}
