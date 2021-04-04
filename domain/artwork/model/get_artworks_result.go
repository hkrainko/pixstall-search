package model

import "pixstall-search/domain/model"

type GetArtworksResult struct {
	Page     model.Page `json:"page"`
	Artworks []Artwork  `json:"artworks"`
}
