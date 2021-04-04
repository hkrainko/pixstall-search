package resp

type Page struct {
	Current     int `json:"current"`
	TotalPage   int `json:"total_pages"`
	TotalResult int `json:"total_results"`
	Size        int `json:"size"`
}