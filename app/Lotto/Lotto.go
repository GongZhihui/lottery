package lotto

import(
    "fmt"
    "lottery/utils"
)

const(
    // MISS .
    MISS = "未中奖";
)

// Manager .
type Manager struct{
    LottoItems map[int]*Item;
    Pos int;
    Notice string;
    Prompt string;
} 

// MakeManager .
func MakeManager() Manager{
    return Manager{make(map[int]*Item), 0, "", ""}
}

// ToString .
func (m *Manager) ToString(){
    for _, it := range m.LottoItems{
        fmt.Println(it.Name, it.Percent, "[", it.Left, ", ", it.Right, "]");
    }
}

// DrawingOnce 抽一次
func (m *Manager) DrawingOnce() (index int, name string){
    var r = utils.MakeRandom(0, int64(Precision));
    for i, it := range m.LottoItems{
        if(it.hasNum(int(r))){
            index = i;
            name = it.Name;
            break;
        }
    }
    return index, name;
}

// CalcMiss 计算未中奖
func (m *Manager) CalcMiss() string {
    var sum = 0.0;
    for _, it := range m.LottoItems{
        sum = sum + it.Percent;
    }
    var miss = 1 - sum;

    var errMsg = "";
    if(miss < 0.0){
        errMsg = "百分比不能大于1!!!";
    }else if(miss > 0.0){
        m.PushItem(MakeItem(MISS, miss));
    }
    return errMsg;
}

// Calcute 计算
func (m *Manager) Calcute() (string){
    // 附加未中奖
    var err = m.CalcMiss();
    if(err == ""){
        MakeItems(m.LottoItems);
    }
    return err;
}

// PushItem add item
func (m *Manager) PushItem(item *Item){
    m.LottoItems[m.Pos] = item;
    var percent = utils.Percent(item.Percent);
    if(float64(percent) > Precision){
        Precision = float64(percent) * 1.0;
    }
    m.Pos = m.Pos + 1;
}

// ClearItem clear item
func (m *Manager) ClearItem(){
    m.Pos = 0
    m.LottoItems = make(map[int]*Item)
}