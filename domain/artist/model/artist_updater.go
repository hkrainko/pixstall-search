package model

import "time"

type ArtistUpdater struct {
	ArtistID string

	UserName    *string
	ProfilePath *string
	State       *UserState

	ArtistIntro    *ArtistIntro
	ArtistBoard    *ArtistBoard
	PaymentMethods *[]string

	CommissionDetails *CommissionDetails
	LastUpdateTime    *time.Time
}
