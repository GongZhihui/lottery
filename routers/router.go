package routers

import (
    "lottery/controllers"
    "github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    
    ns := beego.NewNamespace(
        "/v1",
        beego.NSNamespace("/admin", beego.NSInclude(&controllers.AdminController{})),
        beego.NSNamespace("/drawing", beego.NSInclude(&controllers.DrawingController{})),
    )
    beego.AddNamespace(ns)
}
