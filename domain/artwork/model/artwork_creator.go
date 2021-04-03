package model

import "time"

type ArtworkCreator struct {
	ID               string
	CommissionID     string
	OpenCommissionID string

	ArtistID          string
	ArtistName        string
	ArtistProfilePath *string

	DayUsed   time.Duration
	IsR18     bool
	Anonymous bool

	Path   string
	Rating int

	StartTime     time.Time
	CompletedTime time.Time
}
