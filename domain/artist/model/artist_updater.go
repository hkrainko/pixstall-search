package model

import "time"

type ArtistUpdater struct {
	ArtistID string

	UserName    *string
	ProfilePath *string
	State       *UserState

	// ArtistIntro
	YearOfDrawing *int
	ArtTypes      *[]string

	// ArtistBoard
	Desc *string

	PaymentMethods *[]string

	// CommissionDetails
	CommissionRequestCount *int
	CommissionAcceptCount  *int
	CommissionSuccessCount *int
	AvgRatings             *float64
	LastRequestTime        *time.Time

	LastUpdatedTime *time.Time
}
