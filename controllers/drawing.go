package controllers

import(
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	"strconv"
	"lottery/models"
	"fmt"
	"lottery/app"
	"lottery/utils"
	"time"
)

// DrawingController .
type DrawingController struct {
	beego.Controller
}

// Get .
// @Title Get
// @Description 抽奖
// @Failure 403 submit fail
// @router / [get]
func (c *DrawingController) Get() {
	app.LottoChange = true;
	var lot = app.GetLotto()
	for k, v := range lot.LottoItems{
		var key = "Unit" + strconv.Itoa(k) + "Name"
		c.Data[key] = v.Name
		fmt.Println(key, v.Name)
	}
	c.Data["Prompt"] = lot.Prompt
	c.Data["Notice"] = lot.Notice

	c.TplName = "index.html"
}

// Post .
// @Title Post
// @Description drawing
// @Param	body		body 	models.Item	true  "body for Item content"
// @Success 200 {string} .
// @Failure 403 submit fail
// @router / [post]
func (c *DrawingController) Post() {
	var i, name = app.GetLotto().DrawingOnce()
	var rs = utils.GenerateRandomString(16)
	var prefix = string([]byte(rs[:8]))
	var item = models.Item{Index: i, Name : name, Prefix : prefix}
	c.Data["json"] = item

	var lotInfo models.LotInfo
	lotInfo.ID = utils.GenerateRandomString(5)
	lotInfo.CreateTime = time.Now().Unix()
	lotInfo.Confirm = 0
	lotInfo.Name = name
	lotInfo.Prefix = prefix
	lotInfo.Suffix = rs
	lotInfo.UpdateTime = time.Now().Unix()
	lotInfo.Memo = ""

	models.Insert(&lotInfo)

	c.ServeJSON()
}