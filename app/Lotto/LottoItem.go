package lotto

import(
    //"log"
)

// Precision 中奖率精度 1% =0.01， 0.01% = 0.0001
var Precision = 0.0;

// Item .
type Item struct{
    // item 名字
    Name string;

    // 中奖率
    Percent float64;

    Left int;

    Right int;
}

// MakeItem .
func MakeItem(name string, percent float64)(*Item){
    return &Item{
        Name : name,
        Percent : percent,
        Left : 0,
        Right : 0,
    }
}

// 计算左右区间
func (it *Item) start(left int){
    it.Left = left;
    it.Right = left + int(float64(Precision) * it.Percent) - 1;
}

// 判断num是否在区间
func (it *Item) hasNum(num int) bool{
    return (num >= it.Left) && (num <= it.Right);
}

// MakeItems 计算左右区间
func MakeItems(items map[int]*Item) (bool){
    var lastRight = -1;
    for i := range items{
        items[i].start(lastRight + 1);
        lastRight = items[i].Right;
    }
    return true;
}