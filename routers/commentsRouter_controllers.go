package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["lottery/controllers:AdminController"] = append(beego.GlobalControllerRouter["lottery/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/controllers:AdminController"] = append(beego.GlobalControllerRouter["lottery/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Insert",
            Router: `/insert`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/controllers:AdminController"] = append(beego.GlobalControllerRouter["lottery/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Query",
            Router: `/query`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/controllers:AdminController"] = append(beego.GlobalControllerRouter["lottery/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"get"},
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
