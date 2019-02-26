package app

import (
	"lottery/app/Lotto"
	"lottery/models"
)

var(
	// Lotto .
	lot *lotto.Manager

	// LottoChange 数据库配置是否改变
	LottoChange bool
)

// GetLotto .
func GetLotto() *lotto.Manager {
	if(LottoChange){
		var back = models.MakeLottoBackend()
		back.QueryByID()
		var view = back.ToLottoBackendView()
		lot.Notice = view.Notice
		lot.Prompt = view.LottoPrompt
		lot.ClearItem()
		for _, v := range view.Items{
			lot.PushItem(lotto.MakeItem(v.Name, v.Percent));
		}
		lot.Calcute();
		LottoChange = false
	}
	return lot
}

func init(){
	var lotttManager = lotto.MakeManager()
	lot = &lotttManager	
	LottoChange = false
}

