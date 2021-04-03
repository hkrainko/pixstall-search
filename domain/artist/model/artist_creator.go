package model

type ArtistCreator struct {
	User              `json:",inline"`
	ArtistID          string            `json:"artistId"`
	ArtistIntro       ArtistIntro       `json:"artistIntro"`
	ArtistBoard       ArtistBoard       `json:"artistBoard"`
	PaymentMethods    []string          `json:"paymentMethods"`
	CommissionDetails CommissionDetails `json:"commissionDetails"`
}
