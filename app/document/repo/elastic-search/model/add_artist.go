package model

import (
	"pixstall-search/domain/artist/model"
	"time"
)

type AddArtistRequest struct {
	User                     `json:",inline"`
	ArtistID                 string   `json:"artist_id"`
	ArtistIntroYearOfDrawing int      `json:"year_of_drawing"`
	ArtistIntroArtTypes      []string `json:"art_types"`
	PaymentMethods           []string `json:"payment_methods"`
}

type User struct {
	UserID          string          `json:"user_id"`
	UserName        string          `json:"user_name"`
	ProfilePath     string          `json:"profile_path"`
	State           model.UserState `json:"state"`
	LastUpdatedTime time.Time       `json:"last_updated_time"`
}

func NewAddArtistRequestFromArtistCreator(creator model.ArtistCreator) AddArtistRequest {
	return AddArtistRequest{
		User: User{
			UserID:          creator.UserID,
			UserName:        creator.UserName,
			ProfilePath:     creator.ProfilePath,
			State:           creator.State,
			LastUpdatedTime: creator.LastUpdatedTime,
		},
		ArtistID:                 creator.ArtistID,
		ArtistIntroYearOfDrawing: creator.ArtistIntro.YearOfDrawing,
		ArtistIntroArtTypes:      creator.ArtistIntro.ArtTypes,
		PaymentMethods:           creator.PaymentMethods,
	}
}

type AddArtistResponse struct {
	ID     string   `json:"id"`
	Errors []string `json:"errors"`
}
