package model

type Coin struct {
	Id   int    `json:"id"`
	Coin string `json:"coin"`
}

type CoinItem struct {
	Price float64
	Prec  int // 小数部分的精度
	Time  int64
}
