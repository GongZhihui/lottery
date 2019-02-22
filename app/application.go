package app

import (
	"lottery/app/Lotto"
)

var(
	// Lotto .
	Lotto *lotto.Manager

	// LottoChange 数据库配置是否改变
	LottoChange bool
)

func init(){
	var lotttManager = lotto.MakeManager()
	Lotto = &lotttManager	
	LottoChange = false
}

