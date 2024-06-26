package model

type User struct {
	Id                    int    `json:"id"`
	Username              string `json:"username"`
	Mx_access_key         string `json:"mx_access_key"`
	Mx_secret_key         string `json:"mx_secret_key"`
	Buy_usdt              string `json:"buy_usdt"`
	Min_buy_usdt          string `json:"min_buy_usdt"`
	Profit                string `json:"profit"`
	Loss                  string `json:"loss"`
	Slippage              string `json:"slippage"`
	Next_slippage         string `json:"next_slippage"`
	Worker_count          string `json:"worker_count"`
	Queue_time_gap        string `json:"queue_time_gap"`
	Server_queue_time_gap string `json:"server_queue_time_gap"`
	Request_time_gap      string `json:"request_time_gap"`
	New_request_time_gap  string `json:"new_request_time_gap"`
	Long_profit           string `json:"long_profit"`
	Short_profit          string `json:"short_profit"`
}

type Userpassword struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}
