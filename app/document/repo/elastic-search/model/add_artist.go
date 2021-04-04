package model

import (
	"pixstall-search/domain/artist/model"
	"time"
)

type AddArtistRequest struct {
	ID             string `json:"id"`
	User           `json:",inline"`
	ArtistID       string   `json:"artist_id"`
	PaymentMethods []string `json:"payment_methods"`

	// ArtistIntro
	YearOfDrawing int      `json:"year_of_drawing"`
	ArtTypes      []string `json:"art_types"`

	// ArtistBoard
	Desc string `json:"desc"`

	// CommissionDetails
	CommissionRequestCount int        `json:"commission_request_count"`
	CommissionAcceptCount  int        `json:"commission_accept_count"`
	CommissionSuccessCount int        `json:"commission_success_count"`
	AvgRatings             *int       `json:"avg_ratings"`
	LastRequestTime        *time.Time `json:"last_request_time"`
}

type User struct {
	UserID          string          `json:"user_id"`
	UserName        string          `json:"user_name"`
	ProfilePath     *string         `json:"profile_path"`
	State           model.UserState `json:"state"`
	RegTime         time.Time       `json:"reg_time"`
	LastUpdatedTime time.Time       `json:"last_updated_time"`
}

func NewAddArtistRequest(creator model.ArtistCreator) AddArtistRequest {
	return AddArtistRequest{
		User: User{
			UserID:          creator.UserID,
			UserName:        creator.UserName,
			ProfilePath:     creator.ProfilePath,
			State:           creator.State,
			RegTime:         creator.RegTime,
			LastUpdatedTime: creator.LastUpdatedTime,
		},
		ID:             creator.ArtistID,
		ArtistID:       creator.ArtistID,
		YearOfDrawing:  creator.ArtistIntro.YearOfDrawing,
		ArtTypes:       creator.ArtistIntro.ArtTypes,
		PaymentMethods: creator.PaymentMethods,
	}
}
