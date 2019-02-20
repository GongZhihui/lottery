package models

import(
	"github.com/astaxie/beego/orm"
)

// User user
type User struct {
	Name string `orm:"pk;column(name);"`
	Password string `orm:"column(password);"`
	Role string `orm:"column(role);"`
}

func init() {
	orm.RegisterModel(new(User))
}

// TableName .
func (u *User) TableName() string{
	return "user"
}

// ToString .
func (u *User) ToString() string{
	return "[Name: " + u.Name + ", Password: " + 
		u.Password + ", Role: " + u.Role + "]"
}

// SelectPwdByName .
func (u *User) SelectPwdByName(name string) string{
	var sql = "select * from user where name=?"
	var o = orm.NewOrm()
	o.Using("default")
	if err := o.Raw(sql, name).QueryRow(u); err != nil {
		return err.Error()
	}
	return ""
}
