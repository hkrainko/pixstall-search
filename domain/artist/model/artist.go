package model

import "time"

type Artist struct {
	User
	ArtistID          string            `json:"artistId"`
	ArtistIntro       ArtistIntro       `json:"artistIntro"`
	ArtistBoard       ArtistBoard       `json:"artistBoard"`
	PaymentMethods    []string          `json:"paymentMethods"`
	CommissionDetails CommissionDetails `json:"commissionDetails"`
}

type User struct {
	UserID          string    `json:"userId"`
	UserName        string    `json:"userName"`
	ProfilePath     string    `json:"profilePath"`
	State           UserState `json:"state"`
	RegTime         time.Time `json:"regTime"`
	LastUpdatedTime time.Time `json:"lastUpdatedTime"`
}

type ArtistIntro struct {
	YearOfDrawing int      `json:"yearOfDrawing"`
	ArtTypes      []string `json:"artTypes"`
}

type ArtistBoard struct {
	Desc string `json:"desc" bson:"desc"`
}

type CommissionDetails struct {
	CommissionRequestCount int        `json:"commissionRequestCount"`
	CommissionAcceptCount  int        `json:"commissionAcceptCount"`
	CommissionSuccessCount int        `json:"commissionSuccessCount"`
	AvgRatings             *int       `json:"avgRatings,omitempty"`
	LastRequestTime        *time.Time `json:"lastRequestTime,omitempty"`
}

type UserState string

const (
	UserStatePending    UserState = "P" //Allow to change userID
	UserStateActive     UserState = "A"
	UserStateTerminated UserState = "T"
)
