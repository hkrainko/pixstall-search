package model

type OpenCommissionCreator struct {
	ID               string
	ArtistID         string
	Title            string
	Desc             string
	Price            Price
	DayNeed          DayNeed
	IsR18            bool
	AllowBePrivate   bool
	AllowAnonymous   bool
	SampleImagePaths []string
}
