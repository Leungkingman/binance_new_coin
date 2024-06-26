package request

type MxOrder struct {
	Access_key string	`json:"access_key"`
	Secret_key string	`json:"secret_key"`
	Coin       string	`json:"coin"`
	Price      string	`json:"price"`
	Quantity   string	`json:"quantity"`
}

type Article struct {
	Title string `json:"title"`
}

type BinanceApi struct {

}