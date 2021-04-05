package model

import (
	"pixstall-search/domain/artist/model"
	"time"
)

type UpdateArtistRequest struct {
	ID string `json:"id"`

	UserName    *string          `json:"user_name,omitempty"`
	ProfilePath *string          `json:"profile_path,omitempty"`
	State       *model.UserState `json:"state,omitempty"`

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
	AvgRatings             *float64   `json:"avg_ratings,omitempty"`
	LastRequestTime        *time.Time `json:"last_request_time,omitempty"`
	LastUpdatedTime        *time.Time `json:"last_updated_time,omitempty"`
}

func NewUpdateArtistRequest(updater model.ArtistUpdater) UpdateArtistRequest {
	return UpdateArtistRequest{
		ID:                     updater.ArtistID,
		UserName:               updater.UserName,
		ProfilePath:            updater.ProfilePath,
		State:                  updater.State,
		PaymentMethods:         updater.PaymentMethods,
		YearOfDrawing:          updater.YearOfDrawing,
		ArtTypes:               updater.ArtTypes,
		Desc:                   updater.Desc,
		CommissionRequestCount: updater.CommissionRequestCount,
		CommissionAcceptCount:  updater.CommissionAcceptCount,
		CommissionSuccessCount: updater.CommissionSuccessCount,
		AvgRatings:             updater.AvgRatings,
		LastRequestTime:        updater.LastRequestTime,
		LastUpdatedTime:        updater.LastUpdatedTime,
	}
}
