package main

import (
	_ "lottery/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego/orm"
	//"fmt"
)

func init() {
	beego.InsertFilter("/v1/admin/*", beego.BeforeRouter, func(ctx *context.Context){
		if(ctx.Request.RequestURI == "/v1/admin/login"){
			return;
		}

		var userCookie = ctx.GetCookie("userName")
		var userName = ctx.Input.Session("userName")
		if userName != userCookie{
		 	//ctx.Redirect(302,"/v1/admin/login")
			ctx.Redirect(302, "/")	 
		}
	})

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = false
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	orm.RegisterDriver("sqlite", orm.DRSqlite)
    orm.RegisterDataBase("default", "sqlite3", "./database/lottery.db")
    orm.RunSyncdb("default", false, true)
}

func main() { 
	beego.Run()
}

