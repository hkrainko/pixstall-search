package model

import "time"

type Artwork struct {
	ID               string `json:"id" bson:"id"`
	OpenCommissionID string `json:"openCommissionId" bson:"openCommissionId"`

	ArtistID          string  `json:"artistId" bson:"artistId"`
	ArtistName        string  `json:"artistName" bson:"artistName"`
	ArtistProfilePath *string `json:"artistProfilePath,omitempty" bson:"artistProfilePath,omitempty"`

	DayUsed   time.Duration `json:"dayUsed" bson:"dayUsed"`
	IsR18     bool          `json:"isR18" bson:"isR18"`
	Anonymous bool          `json:"anonymous" bson:"anonymous"`

	Path   string `json:"path" bson:"path"`
	Rating int    `json:"rating" bson:"rating"`

	Title       string          `json:"title"`
	TextContent string          `json:"textContent"`
	Views       int             `json:"views"`
	FavorCount  int             `json:"favorCount"`

	CreateTime     time.Time    `json:"createTime" bson:"createTime"`
	StartTime      time.Time    `json:"startTime" bson:"startTime"`
	CompletedTime  time.Time    `json:"completedTime" bson:"completedTime"`
	LastUpdateTime time.Time    `json:"lastUpdateTime" bson:"lastUpdateTime"`
	State          ArtworkState `json:"state" bson:"state"`
}

type ArtworkState string

const (
	ArtworkStateActive  ArtworkState = "Active"
	ArtworkStateHidden  ArtworkState = "Hidden"
	ArtworkStateDeleted ArtworkState = "Deleted"
)
