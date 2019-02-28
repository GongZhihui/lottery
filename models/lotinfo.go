package models

import(
    "github.com/astaxie/beego/orm"
)

// LotInfo lot_info
type LotInfo struct {
    ID string `orm:"pk;column(id);"`
    Name string `orm:"column(name);"`
    Prefix string `orm:"column(prefix);"`
    // suffix其实是整个字符串不是后面一部分
    Suffix string `orm:"column(suffix);"`
    Confirm int `orm:"column(confirm);"`
    CreateTime int64 `orm:"column(create_time);"`
    UpdateTime int64 `orm:"column(update_time);"`
    Memo string `orm:"column(memo);"`
}

func init() {
    orm.RegisterModel(new(LotInfo)) 
}

// TableName .
func (u *LotInfo) TableName() string{
    return "lot_info"
}

// QueryByPrefix .
func (u *LotInfo) QueryByPrefix() string{
    var sql = "select * from lot_info where prefix=?"
    var o = orm.NewOrm()
    o.Using("default")
    if err := o.Raw(sql, u.Prefix).QueryRow(u); err != nil {
        return err.Error()
    }
    return ""
}
