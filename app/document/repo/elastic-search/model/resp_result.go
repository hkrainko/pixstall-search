package model

type ResponseResult struct {
	ID     string   `json:"id"`
	Errors []string `json:"errors"`
}
