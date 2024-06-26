package model

type LogData struct {
	Success     bool
	StatusCode  int
	Time        int64
	Content     string
	Level       string // normal——正常输出币安内容；new——发现新币种；error——错误；success——下单成功
	Source      string // announce——公告；newListing——新币公告；mobileApi——手机端api；exchange——交易所；others——其他
	DomainIndex int
	CreateTime  int64
}

type NewLogData struct {
	Success     bool
	StatusCode  int
	Time        int64
	Content     string
	TitleType   string
	PublishTime string // 新标题发布时间，从接口获取
	Level       string // normal——正常输出币安内容；new——发现新币种；error——错误；success——下单成功
	Source      string // announce——公告；newListing——新币公告；mobileApi——手机端api；exchange——交易所；others——其他
	Domain      string
}
