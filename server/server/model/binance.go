package model

type AnnounceApiResponse struct {
	Code          int64       `json:"code"`
	Message       string      `json:"message"`
	MessageDetail string      `json:"messageDetail"`
	Success       bool        `json:"success"`
	Data          interface{} `json:"data"`
}

type AllCoin struct {
	Data []CoinItem `json:"data"`
}

type CrawlerRequest struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Items []Announcement
}

type Announcement struct {
	title string `json:"title"`
}
