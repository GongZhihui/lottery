package main

import (
	_ "lottery/routers"
	"github.com/astaxie/beego"

	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego/orm"
)

func init1() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
    orm.RegisterDataBase("default", "sqlite3", "./database/lottery.db")
    orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}

