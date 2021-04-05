package msg

import (
	"pixstall-search/domain/artist/model"
	"time"
)

type UpdatedArtist struct {
	ArtistID string `json:"artistId"`

	UserName    *string          `json:"userName,omitempty"`
	ProfilePath *string          `json:"profilePath,omitempty"`
	Email       *string          `json:"email,omitempty"`
	Birthday    *string          `json:"birthday,omitempty"`
	Gender      *string          `json:"gender,omitempty"`
	State       *model.UserState `json:"state,omitempty"`

	PaymentMethods *[]string `json:"paymentMethods,omitempty"`

	// ArtistIntro
	YearOfDrawing *int      `json:"yearOfDrawing"`
	ArtTypes      *[]string `json:"artTypes"`

	// ArtistBoard
	BannerPath *string `json:"bannerPath"`
	Desc       *string `json:"desc"`

	// CommissionDetails
	CommissionRequestCount *int       `json:"commissionRequestCount,omitempty"`
	CommissionAcceptCount  *int       `json:"commissionAcceptCount,omitempty"`
	CommissionSuccessCount *int       `json:"commissionSuccessCount,omitempty"`
	AvgRatings             *float64   `json:"avgRatings,omitempty"`
	LastRequestTime        *time.Time `json:"lastRequestTime,omitempty"`
	LastUpdatedTime        *time.Time `json:"lastUpdatedTime,omitempty"`
}

func (u UpdatedArtist) ToDomainArtistUpdater() model.ArtistUpdater {
	return model.ArtistUpdater{
		ArtistID:               u.ArtistID,
		UserName:               u.UserName,
		ProfilePath:            u.ProfilePath,
		State:                  u.State,
		YearOfDrawing:          u.YearOfDrawing,
		ArtTypes:               u.ArtTypes,
		Desc:                   u.Desc,
		PaymentMethods:         u.PaymentMethods,
		CommissionRequestCount: u.CommissionRequestCount,
		CommissionAcceptCount:  u.CommissionAcceptCount,
		CommissionSuccessCount: u.CommissionSuccessCount,
		AvgRatings:             u.AvgRatings,
		LastRequestTime:        u.LastRequestTime,
		LastUpdatedTime:        u.LastUpdatedTime,
	}
}
