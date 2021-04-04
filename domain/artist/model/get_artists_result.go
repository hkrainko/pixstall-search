package model

import "pixstall-search/domain/model"

type GetArtistsResult struct {
	Page    model.Page `json:"page"`
	Artists []Artist   `json:"artists"`
}
