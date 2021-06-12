package search_artworks

import (
	"pixstall-search/domain/artwork/model"
	model2 "pixstall-search/domain/model"
)

type Response struct {
	Page     model2.Page     `json:"page"`
	Artworks []model.Artwork `json:"artworks"`
}

func NewResponse(result model.GetArtworksResult) *Response {
	return &Response{
		Page:    result.Page,
		Artworks: result.Artworks,
	}
}
