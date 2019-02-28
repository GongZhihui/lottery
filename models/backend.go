package models

import(
    "github.com/astaxie/beego/orm"
    "encoding/json"
    //"fmt"
)

// LottoBackend LottoBackend
type LottoBackend struct {
    ID string `orm:"pk;default(4016);column(id);"`
    Notice string `orm:"column(notice);"`
    LottoPrompt string `orm:"column(lotto_prompt);"`
    Items string `orm:"column(items)"`
}

// Item .
type Item struct{
    Index int `json:"index"`
    Name string `json:"name"`
    Percent float64 `json:"percent"`
    Prefix string `json:"prefix"`
}

// SerializeItem .
func (v *LottoBackendView) SerializeItem() string{
    var js, _ = json.Marshal(v.Items)
    return string(js[:])
}

// ParseItem .
func (v *LottoBackendView) ParseItem(s string) {
    json.Unmarshal([]byte(s), &v.Items)
}

// LottoBackendView .
type LottoBackendView struct {
    Notice string `json:"notice"`
    LottoPrompt string `json:"lottoPrompt"`
    Items []Item `json:"items"`
}

// MakeLottoBackendView . 
func MakeLottoBackendView() LottoBackendView {
    return LottoBackendView{"","", make([]Item, 0)}
}

func init() {
    orm.RegisterModel(new(LottoBackend))
}

// MakeLottoBackend .
func MakeLottoBackend() LottoBackend{
    return LottoBackend{ID:"4016"}
}

// TableName .
func (l *LottoBackend) TableName() string{
    return "lotto_backend"
}

// ToString .
func (l *LottoBackend) ToString() string{
    return "[ID:" + l.ID + 
        ", Notice:" + l.Notice + 
        ", LottoPrompt:" + l.LottoPrompt +
        ", Items:" + l.Items + "]"
}

// ToLottoBackendView .
func (l *LottoBackend) ToLottoBackendView() LottoBackendView{
    var view = MakeLottoBackendView()
    view.ParseItem(l.Items)
    view.Notice = l.Notice
    view.LottoPrompt = l.LottoPrompt
    return view
}

// QueryByID .
func (l *LottoBackend) QueryByID() string { 
    var o = orm.NewOrm()
    o.Using("default")
    if err := o.Read(l); err != nil {
        return err.Error()
    }
    return ""
}

