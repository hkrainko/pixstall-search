package model

import "time"

type ArtworkUpdater struct {
	ID                   string  `json:"id"`
	ArtistName           *string `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath"`

	Title       *string `json:"title"`
	TextContent *string `json:"textContent"`
	Views       *int     `json:"views"`
	FavorCount  *int    `json:"favorCount"`

	LastUpdatedTime *time.Time `json:"lastUpdatedTime" bson:"lastUpdatedTime"`
	State           *ArtworkState
}
