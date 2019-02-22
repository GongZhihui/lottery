package models

import(
	"github.com/astaxie/beego/orm"
	"lottery/app/Lotto"
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
	Name string `json:"name"`
	Percent float64 `json:"percent"`
}

// LottoBackendView .
type LottoBackendView struct {
	Notice string `json:"notice"`
	LottoPrompt string `json:"lottoPrompt"`
	Items []Item `json:"items"`
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
	var lottItems = lotto.ParseItems(l.Items)
	var items = make([]Item, 0)
	for _, v := range lottItems {
		items = append(items, Item{v.Name, v.Percent})
	}
	return LottoBackendView{l.Notice,
		l.LottoPrompt, items}
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
