package models

import(
	"github.com/astaxie/beego/orm"
)

// Insert .
func Insert(object interface {}) (string) {
	var o = orm.NewOrm()
	o.Using("default")
	if _, err := o.Insert(object); err != nil {
		return err.Error()
	}
	return ""
}

// Update .
func Update(object interface {}) (string) {
	var o = orm.NewOrm()
	o.Using("default")
	if _, err := o.Update(object); err != nil {
		return err.Error()
	}
	return ""
}

// Delete .
func Delete(object interface {}) (string) {
	var o = orm.NewOrm()
	o.Using("default")
	if _, err := o.Delete(object); err != nil {
		return err.Error()
	}
	return ""
}