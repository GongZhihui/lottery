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
func Update(object interface {}, cols ...string) (string) {
    var o = orm.NewOrm()
    o.Using("default")
    if _, err := o.Update(object, cols...); err != nil {
        return err.Error()
    }
    return ""
}

// Delete .
func Delete(object interface {}, cols ...string) (string) {
    var o = orm.NewOrm()
    o.Using("default")
    if _, err := o.Delete(object, cols...); err != nil {
        return err.Error()
    }
    return ""
}