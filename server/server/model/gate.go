package model

type GateCurrenciesItem struct {
	Currency          string `json:"currency"`
	Delisted          bool   `json:"delisted"`
	Withdraw_disabled bool   `json:"withdraw_disabled"`
	Withdraw_delayed  bool   `json:"withdraw_delayed"`
	Deposit_disabled  bool   `json:"deposit_disabled"`
	Trade_disabled    bool   `json:"trade_disabled"`
}

type GateTickerItem struct {
	Currency_pair string `json:"currency_pair"`
	Last          string `json:"last"`
	Lowest_ask    string `json:"lowest_ask"`
	Highest_bid   string `json:"highest_bid"`
	Base_volume   string `json:"base_volume"`
	Quote_volume  string `json:"quote_volume"`
}

type TestGateOrder struct {
	Coin        string `json:"coin"`
	Usdt        string `json:"usdt"`
	Slippage    string `json:"slippage"`
	Price       string `json:"price"`
	Side        string `json:"side"`
	Account     string `json:"account"`
	Auto_borrow bool   `json:"auto_borrow"`
}

type GateOrderResult struct {
	Id            string `json:"id"`
	Text          string `json:"text"`
	Create_time   string `json:"create_time"`
	Update_time   string `json:"update_time"`
	Status        string `json:"status"`
	Currency_pair string `json:"currency_pair"`
	Type          string `json:"type"`
	Account       string `json:"account"`
	Side          string `json:"side"`
	Amount        string `json:"amount"`
	Price         string `json:"price"`
	Time_in_force string `json:"time_in_force"`
	Iceberg       string `json:"iceberg"`
	Left          string `json:"left"`
	Filled_total  string `json:"filled_total"`
	Fee           string `json:"fee"`
	Fee_currency  string `json:"fee_currency"`
}

type DB_Gate_Order struct {
	Id          uint   `json:"id" gorm:"primarykey"`
	Order_id    string `json:"order_id"`
	Coin        string `json:"coin"`
	Price       string `json:"price"`
	Find_price  string `json:"find_price"`
	Find_time   string `json:"find_time"`
	Amount      string `json:"amount"`
	Total       string `json:"total"`
	Left        string `json:"left"`
	Create_time string `json:"create_time"`
}

type TransfersBody struct {
	Currency      string `json:"currency"`
	From          string `json:"from"`
	To            string `json:"to"`
	Amount        string `json:"amount"`
	Currency_pair string `json:"currency_pair"`
}
