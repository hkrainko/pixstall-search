package model

type ArtistCreator struct {
	User              `json:",inline"`
	ArtistID          string            `json:"artistId"`
	ArtistIntro       ArtistIntro       `json:"artistIntro"`
	PaymentMethods    []string          `json:"paymentMethods"`
}
