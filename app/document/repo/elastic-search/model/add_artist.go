package model

import (
	"pixstall-search/domain/artist/model"
	"time"
)

type AddArtistRequest struct {
	ID             string `json:"id"`
	User           `json:",inline"`
	ArtTypes       []string `json:"art_types"`
	PaymentMethods []string `json:"payment_methods"`

	// ArtistIntro
	ArtistID      string `json:"artist_id"`
	YearOfDrawing int    `json:"year_of_drawing"`

	// CommissionDetails
	CommissionRequestCount int        `json:"commission_request_count"`
	CommissionAcceptCount  int        `json:"commission_accept_count"`
	CommissionSuccessCount int        `json:"commission_success_count"`
	AvgRatings             *int       `json:"avg_ratings,omitempty"`
	LastRequestTime        *time.Time `json:"last_request_time,omitempty"`
}

type User struct {
	UserID          string          `json:"user_id"`
	UserName        string          `json:"user_name"`
	ProfilePath     string          `json:"profile_path"`
	State           model.UserState `json:"state"`
	RegTime         time.Time       `json:"reg_time"`
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
		ID:             creator.ArtistID,
		ArtistID:       creator.ArtistID,
		YearOfDrawing:  creator.ArtistIntro.YearOfDrawing,
		ArtTypes:       creator.ArtistIntro.ArtTypes,
		PaymentMethods: creator.PaymentMethods,
	}
}

type AddArtistResponse struct {
	ID     string   `json:"id"`
	Errors []string `json:"errors"`
}
