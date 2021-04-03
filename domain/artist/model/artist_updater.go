package model

type ArtistUpdater struct {
	ArtistID string

	UserName    *string
	ProfilePath *string
	Email       *string
	Birthday    *string
	Gender      *string
	State       *string
	RegTime     *string

	ArtistIntro    *ArtistIntro
	ArtistBoard    *ArtistBoard
	PaymentMethods *[]string
}
