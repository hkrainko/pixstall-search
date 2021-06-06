package model

type Page struct {
	Current     int `json:"current"`
	TotalPage   int `json:"totalPage"`
	TotalResult int `json:"totalResult"`
	Size        int `json:"size"`
}
