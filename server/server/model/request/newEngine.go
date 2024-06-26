package request

type SingleServerFetchTask struct {
	Domains  []string `json:"domains"`
	Titles   []string `json:"titles"`
	Coins    []string `json:"coins"`
	BuyUsdt  string   `json:"buyUsdt"`
	Slippage string   `json:"slippage"`
}
