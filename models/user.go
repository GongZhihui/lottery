package models

import(

)

// User user
type User struct {
	id int `orm:"pk;column(id);"`
	name string `orm:"column(name);"`
	password string `orm:"column(password);"`
	role string `orm:"column(role);"`
}