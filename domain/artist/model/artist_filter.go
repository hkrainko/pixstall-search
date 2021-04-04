package model

import "time"

type ArtistFilter struct {
	State *[]UserState

	// RegTime
	RegTimeFrom *time.Time
	RegTimeTo   *time.Time

	PaymentMethods *[]string

	// YearOfDrawing
	YearOfDrawingFrom *int
	YearOfDrawingTo   *int

	// CommissionRequestCount
	CommissionRequestCountFrom *int
	CommissionRequestCountTo   *int

	// CommissionSuccessCount
	CommissionSuccessCountFrom *int
	CommissionSuccessCountTo   *int

	// AvgRatings
	AvgRatingsMin *int

	// LastRequestTime
	LastRequestTimeFrom *time.Time
	LastRequestTimeTo   *time.Time
}
