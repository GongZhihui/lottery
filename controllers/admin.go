package controllers

import(
    "github.com/astaxie/beego"
    //"strconv"
    "encoding/json"
    "lottery/models"
    "lottery/app"
    "fmt"
    "time"
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
    c.Data["json"] = view
    c.ServeJSON()
}

// Notice .
// @Title Notice
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
// @Title Prompt
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
// @Title LottoItem
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

    // 通知item改变
    app.LottoChange = true
    c.ServeJSON()
}

// LotInfo .
// @Title LotInfo
// @Description 提交中奖信息
// @Param	body		body 	models.LotInfo	true  "body for LotInfo.Prefix content"
// @Success 200 {object} models.LotInfo
// @Failure 403 submit fail
// @router /lotinfo [post]
func (c *AdminController) LotInfo() {
    var info models.LotInfo
    json.Unmarshal(c.Ctx.Input.RequestBody, &info)
 
    if err := info.QueryByPrefix(); err == "" {
        c.Data["json"] = info
    }else{
        c.Data["json"] = ""
    }
    c.ServeJSON()
}

// LotInfoUpdate .
// @Title LotInfoUpdate
// @Description 修改中奖信息
// @Param	body		body 	models.LotInfo	true  "body for LotInfo.Prefix content"
// @Success 200 {bool} true
// @Failure 403 submit fail
// @router /lotinfo_update [post]
func (c *AdminController) LotInfoUpdate() {
    var info models.LotInfo
    json.Unmarshal(c.Ctx.Input.RequestBody, &info)
    info.UpdateTime = time.Now().Unix()
    models.Update(&info, "Confirm", "Memo", "UpdateTime")
    c.Data["json"] = true
    c.ServeJSON()
}

// LoginInit .
// @Title LoginInit
// @Description 登录
// @Failure 403 submit fail
// @router /login [get]
func (c *AdminController) LoginInit() {
    c.TplName = "admin/login.html"
}

// Main .
// @Title Main
// @Description 登录
// @Failure 403 submit fail
// @router /main [get]
func (c *AdminController) Main() {
    c.TplName = "admin/admin.html"
}

// Login .
// @Title Login
// @Description 登录
// @Param	body		body 	models.User	true  "body for User content"
// @Success 200 {string} 跳转到main
// @Failure 403 submit fail
// @router /login [post]
func (c *AdminController) Login() {
    var user models.User
    var name = c.GetString("name")
    var pwd = c.GetString("password")

    fmt.Println(name, pwd)

    user.SelectPwdByName(name)
    var lhs = user.Password
    var rhs = user.EncodePassword(pwd)
    fmt.Println(lhs, rhs)
    if(lhs == rhs) {
        c.SetSession("userName", name)
        c.Ctx.SetCookie("userName", name, 30, "/v1/")
        c.Ctx.Redirect(302, "/v1/admin/main")
    }else{
        c.Ctx.Redirect(302, "/v1/admin/login")
    }
}