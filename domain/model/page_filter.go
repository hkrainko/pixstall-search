package model

type PageFilter struct {
	Current int `json:"current"` // Current page
	Size    int `json:"size"`    // Size per page
}
