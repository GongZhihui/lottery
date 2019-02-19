package models

import(
	"github.com/astaxie/beego/orm"
)

// LotInfo lot_info
type LotInfo struct {
	ID string `orm:"pk;column(id);"`
	Name string `orm:"column(name);"`
	Prefix string `orm:"column(prefix);"`
	Suffix string `orm:"column(suffix);"`
	Confirm int `orm:"column(confirm);"`
	CreateTime int64 `orm:"column(create_time);"`
	UpdateTime int64 `orm:"column(update_time);"`
	Memo string `orm:"column(memo);"`
}

func init() {
	orm.RegisterModel(new(LotInfo)) 
}