package model

import (
	model4 "pixstall-search/domain/model"
	"time"
)

type ArtistFilter struct {
	State *[]UserState

	// RegTime
	RegTime *TimeRange

	PaymentMethods *[]string

	// YearOfDrawing
	YearOfDrawing *TimeRange

	// CommissionRequestCount
	CommissionRequestCount *IntRange

	// CommissionRequestCount
	CommissionAcceptCount *IntRange

	// CommissionSuccessCount
	CommissionSuccessCount *IntRange

	// AvgRatings
	AvgRatings *FloatRange

	// LastRequestTime
	LastRequestTime *TimeRange

	PageFilter      model4.PageFilter
}

type TimeRange struct {
	From *time.Time `json:"from,omitempty"`
	To   *time.Time `json:"to,omitempty"`
}

type IntRange struct {
	From *time.Time `json:"from,omitempty"`
	To   *time.Time `json:"to,omitempty"`
}

type FloatRange struct {
	From *float64 `json:"from,omitempty"`
	To   *float64 `json:"to,omitempty"`
}
