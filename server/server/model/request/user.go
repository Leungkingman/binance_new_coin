package request

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PasswordStruct struct {
	Password string `json:"password"`
}

type UserInfoStruct struct {
	Buy_usdt              string `json:"buy_usdt"`
	Min_buy_usdt          string `json:"min_buy_usdt"`
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

type UserSecretStruct struct {
	Mx_access_key string `json:"mx_access_key"`
	Mx_secret_key string `json:"mx_secret_key"`
}
