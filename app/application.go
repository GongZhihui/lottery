package app

import (
	"lottery/app/Lotto"
)

var(
	// Lotto .
	Lotto *lotto.Manager
)

func init(){
	var lotttManager = lotto.MakeManager()
	Lotto = &lotttManager	
}

