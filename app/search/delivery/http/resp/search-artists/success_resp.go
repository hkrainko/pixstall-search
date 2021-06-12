package search_artists

import (
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/model"
)

type Response struct {
	Page    model2.Page    `json:"page"`
	Artists []model.Artist `json:"artists"`
}

func NewResponse(result model.GetArtistsResult) *Response {
	return &Response{
		Page:    result.Page,
		Artists: result.Artists,
	}
}