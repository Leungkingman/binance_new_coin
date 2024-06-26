package model

type MxOrder struct {
	Order_id    string
	Create_time string
}

type DB_Order_Item struct {
	Id       int
	Order_id string
}

type DB_Order struct {
	Data []DB_Order_Item `json:"data"`
}

type MxOrderItem struct {
	Id            string `json:"id"`
	Symbol        string `json:"symbol"`
	Price         string `json:"price"`
	Quantity      string `json:"quantity"`
	State         string `json:"state"`
	Type          string `json:"type"`
	Deal_quantity string `json:"deal_quantity"`
	Deal_amount   string `json:"deal_amount"`
	Create_time   int64  `json:"create_time"`
}

type MxSymbolsItem struct {
	Symbol         string `json:"symbol"`
	Price_scale    int    `json:"price_scale"`
	Quantity_scale int    `json:"quantity_scale"`
	Min_amount     string `json:"min_amount"`
	Max_amount     string `json:"max_amount"`
	Maker_fee_rate string `json:"maker_fee_rate"`
	Taker_fee_rate string `json:"taker_fee_rate"`
	State          string `json:"state"`
}

type MxOrderResult struct {
	Code int           `json:"code"`
	Data []MxOrderItem `json:"data"`
}

type MxSymbolResult struct {
	Code int             `json:"code"`
	Data []MxSymbolsItem `json:"data"`
}

type CoinPrice struct {
	Coin  string `json:"coin"`
	Price string `json:"price"`
}

type CoinError struct {
	Coin string
	Err  error
}
