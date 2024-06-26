package model

type Article struct {
	Id          uint   `json:"id" gorm:"primarykey"`
	Content     string `json:"content"`
	Create_time string `json:"create_time"`
}

type ApiArticle struct {
	Title string `json:"title"`
	ReleaseDate string `json:"releaseDate"`
}

type Notice struct {
	Id           uint   `json:"id" gorm:"primarykey"`
	Content      string `json:"content"`
	Create_time  string `json:"create_time"`
	Title_type   string `json:"title_type"`
	Publish_time string `json:"publish_time"`
}

type News struct {
	Id    int
	Title string
}
