package model

import (
	"pixstall-search/domain/artist/model"
	"time"
)

type AddArtistRequest struct {
	User           `json:",inline"`
	ArtistID       string      `json:"artist_id"`
	ArtistIntro    ArtistIntro `json:"artist_intro"`
	PaymentMethods []string    `json:"payment_methods"`
}

type User struct {
	UserID          string          `json:"user_id"`
	UserName        string          `json:"user_name"`
	ProfilePath     string          `json:"profile_path"`
	State           model.UserState `json:"state"`
	LastUpdatedTime time.Time       `json:"last_updated_time"`
}

type ArtistIntro struct {
	YearOfDrawing int      `json:"year_Of_drawing"`
	ArtTypes      []string `json:"art_types"`
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
		ArtistID: creator.ArtistID,
		ArtistIntro: ArtistIntro{
			YearOfDrawing: creator.ArtistIntro.YearOfDrawing,
			ArtTypes:      creator.ArtistIntro.ArtTypes,
		},
		PaymentMethods: creator.PaymentMethods,
	}
}

type AddArtistResponse struct {
	ID     string      `json:"id"`
	Errors interface{} `json:"errors"`
}
