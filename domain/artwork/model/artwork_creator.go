package model

import "time"

type ArtworkCreator struct {
	CommissionID     string `json:"commissionId"`
	OpenCommissionID string `json:"openCommissionId"`

	ArtistID             string  `json:"artistId"`
	ArtistName           string  `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath,omitempty"`

	DayUsed   time.Duration `json:"dayUsed"`
	IsR18     bool          `json:"isR18"`
	Anonymous bool          `json:"anonymous"`

	Path               string  `json:"path"`
	Rating             int     `json:"rating"`

	StartTime     time.Time `json:"startTime"`
	CompletedTime time.Time `json:"completedTime"`
}
