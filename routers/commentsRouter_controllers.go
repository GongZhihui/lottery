package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["lottery/controllers:AdminController"] = append(beego.GlobalControllerRouter["lottery/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/controllers:AdminController"] = append(beego.GlobalControllerRouter["lottery/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Check",
            Router: `/check`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/controllers:AdminController"] = append(beego.GlobalControllerRouter["lottery/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/controllers:AdminController"] = append(beego.GlobalControllerRouter["lottery/controllers:AdminController"],
        beego.ControllerComments{
            Method: "LoginInit",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/controllers:AdminController"] = append(beego.GlobalControllerRouter["lottery/controllers:AdminController"],
        beego.ControllerComments{
            Method: "LottoItem",
            Router: `/lotto_item`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/controllers:AdminController"] = append(beego.GlobalControllerRouter["lottery/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Main",
            Router: `/main`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/controllers:AdminController"] = append(beego.GlobalControllerRouter["lottery/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Notice",
            Router: `/notice`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/controllers:AdminController"] = append(beego.GlobalControllerRouter["lottery/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Prompt",
            Router: `/prompt`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/controllers:DrawingController"] = append(beego.GlobalControllerRouter["lottery/controllers:DrawingController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
