package model

import "time"

type OpenCommissionUpdater struct {
	ID               string
	Title            *string
	Desc             *string
	Price            *Price
	DayNeed          *DayNeed
	SampleImagePaths *[]string
	IsR18            *bool
	AllowBePrivate   *bool
	AllowAnonymous   *bool
	State            *OpenCommissionState
	LastUpdatedTime  time.Time
}
