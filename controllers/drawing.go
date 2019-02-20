package controllers

import(
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	//"strconv"

	//"lottery/models"
	"fmt"
	"lottery/app"
	"lottery/app/Lotto"
)

// DrawingController .
type DrawingController struct {
	beego.Controller
}

// @router / [get]
func (c *DrawingController) Get() {
	var m = app.Lotto
	m.PushItem(lotto.MakeItem("A", 0.1));
	m.PushItem(lotto.MakeItem("B", 0.01));
	m.PushItem(lotto.MakeItem("C", 0.0002));
	m.PushItem(lotto.MakeItem("D", 0.15));
	m.PushItem(lotto.MakeItem("E", 0.06));
	m.Calcute();
	m.ToString();
	var times = make(map[string]int, 0);
	var max = 100;
	for i := 0; i< max ; i++{
		var _, name = m.DrawingOnce()
		times[name] = times[name] + 1;
	}

	fmt.Printf("\n");

	fmt.Println("--模拟抽奖--", max , "次");

	for k, v := range times{
		fmt.Printf("%s, 出现 %d 次, 概率为：%f\n", k, v, float64(v)/float64(max));
	}
	c.Ctx.WriteString("content string")
}