package model

type ArtistUpdater struct {
	ArtistID string

	UserName    *string
	ProfilePath *string
	State       *string
	RegTime     *string

	ArtistIntro    *ArtistIntro
	ArtistBoard    *ArtistBoard
	PaymentMethods *[]string

	CommissionDetails *CommissionDetails
}
