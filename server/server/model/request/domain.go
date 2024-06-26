package request

type AddDomain struct {
	Domain string `json:"domain"`
}

type UpdateDomain struct {
	Id     int    `json:"id"`
	Domain string `json:"domain"`
}

type DeleteDomain struct {
	Id     int    `json:"id"`
}
