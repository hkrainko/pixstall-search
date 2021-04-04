package model

import (
	"pixstall-search/domain/model"
)

type ArtistFilter struct {
	State *[]UserState

	// RegTime
	RegTime *model.TimeRange

	PaymentMethods *[]string

	// YearOfDrawing
	YearOfDrawing *model.TimeRange

	// CommissionRequestCount
	CommissionRequestCount *model.IntRange

	// CommissionRequestCount
	CommissionAcceptCount *model.IntRange

	// CommissionSuccessCount
	CommissionSuccessCount *model.IntRange

	// AvgRatings
	AvgRatings *model.FloatRange

	// LastRequestTime
	LastRequestTime *model.TimeRange

	PageFilter model.PageFilter
}
