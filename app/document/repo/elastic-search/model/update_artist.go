package model

import (
	"pixstall-search/domain/artist/model"
	"time"
)

type UpdateArtistRequest struct {
	ID string `json:"id"`

	UserName        *string          `json:"user_name,omitempty"`
	ProfilePath     *string          `json:"profile_path,omitempty"`
	State           *model.UserState `json:"state,omitempty"`
	LastUpdatedTime *time.Time       `json:"last_updated_time,omitempty"`

	PaymentMethods *[]string `json:"payment_methods,omitempty"`

	// ArtistIntro
	YearOfDrawing *int      `json:"year_of_drawing,omitempty"`
	ArtTypes      *[]string `json:"art_types,omitempty"`

	//ArtistBoard
	Desc *string `json:"desc"`

	// CommissionDetails
	CommissionRequestCount *int       `json:"commission_request_count,omitempty"`
	CommissionAcceptCount  *int       `json:"commission_accept_count,omitempty"`
	CommissionSuccessCount *int       `json:"commission_success_count,omitempty"`
	AvgRatings             *int       `json:"avg_ratings,omitempty"`
	LastRequestTime        *time.Time `json:"last_request_time,omitempty"`
}

func NewUpdateArtistRequest(updater model.ArtistUpdater) UpdateArtistRequest {
	var yearOfDrawing *int
	var artTypes *[]string
	if updater.ArtistIntro != nil {
		yearOfDrawing = &updater.ArtistIntro.YearOfDrawing
		artTypes = &updater.ArtistIntro.ArtTypes
	}

	return UpdateArtistRequest{
		ID:                     updater.ArtistID,
		UserName:               updater.UserName,
		ProfilePath:            updater.ProfilePath,
		State:                  updater.State,
		LastUpdatedTime:        updater.LastUpdateTime,
		PaymentMethods:         updater.PaymentMethods,
		YearOfDrawing:          yearOfDrawing,
		ArtTypes:               artTypes,
		Desc:                   &updater.ArtistBoard.Desc,
		CommissionRequestCount: &updater.CommissionDetails.CommissionRequestCount,
		CommissionAcceptCount:  &updater.CommissionDetails.CommissionAcceptCount,
		CommissionSuccessCount: &updater.CommissionDetails.CommissionSuccessCount,
		AvgRatings:             updater.CommissionDetails.AvgRatings,
		LastRequestTime:        updater.CommissionDetails.LastRequestTime,
	}
}

type UpdateArtistResponse struct {
	ID     string   `json:"id"`
	Errors []string `json:"errors"`
}
