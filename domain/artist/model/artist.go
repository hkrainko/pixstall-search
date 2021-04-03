package model

import "time"

type Artist struct {
	ArtistID        string    `json:"artistId"`
	UserName        string    `json:"userName"`
	ProfilePath     string    `json:"profilePath"`
	YearOfDrawing   int       `json:"yearOfDrawing"`
	ArtTypes        []string  `json:"artTypes"`
	Desc            string    `json:"desc" bson:"desc"`
	State           UserState `json:"state"`
	RegTime         time.Time `json:"regTime"`
	LastUpdatedTime time.Time `json:"lastUpdatedTime"`
}

type UserState string

const (
	UserStatePending    UserState = "P" //Allow to change userID
	UserStateActive     UserState = "A"
	UserStateTerminated UserState = "T"
)
