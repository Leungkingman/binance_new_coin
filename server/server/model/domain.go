package model

type DB_Domain struct {
	Id     int    `json:"id" gorm:"primarykey"`
	Domain string `json:"domain"`
}

type DomainInConcurrent struct {
	Domain     string
	Status     int
	Error_time int64
}
