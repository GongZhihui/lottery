package models

import(
	"github.com/astaxie/beego/orm"
)

// User user
type User struct {
	ID string `orm:"pk;column(id);"`
	Name string `orm:"column(name);"`
	Password string `orm:"column(password);"`
	Role string `orm:"column(role);"`
}

func init() {
	orm.RegisterModel(new(User))
}

// TableName .
func (u *User) TableName() string{
	return "table"
}

// ToString .
func (u *User) ToString() string{
	return "[ID: " + u.ID + ", Name: " + u.Name + ", Password: " + 
		u.Password + ", Role: " + u.Role + "]"
}