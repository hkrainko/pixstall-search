package model

import "time"

type ArtworkUpdater struct {
	ID                   string  `json:"id"`
	ArtistName           *string `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath"`

	Title       *string `json:"title"`
	TextContext *string `json:"textContext"`
	Views       *int     `json:"views"`
	FavorCount  *int    `json:"favorCount"`

	LastUpdateTime *time.Time `json:"lastUpdateTime" bson:"lastUpdateTime"`
	State          *ArtworkState
}
