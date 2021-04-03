package msg

import (
	"pixstall-search/domain/artwork/model"
	"time"
)

type CreatedArtwork struct {
	ID                   string             `json:"id" bson:"id"`
	CommissionID         string             `json:"commissionId"`
	OpenCommissionID     string             `json:"openCommissionId"`
	ArtistID             string             `json:"artistId"`
	ArtistName           string             `json:"artistName"`
	ArtistProfilePath    *string            `json:"artistProfilePath,omitempty"`
	RequesterID          string             `json:"requesterId"`
	RequesterName        string             `json:"requesterName"`
	RequesterProfilePath *string            `json:"requesterProfilePath,omitempty"`
	DayUsed              time.Duration      `json:"dayUsed"`
	IsR18                bool               `json:"isR18"`
	Anonymous            bool               `json:"anonymous"`
	Path                 string             `json:"path"`
	Volume               int64              `json:"volume"`
	Size                 Size               `json:"size"`
	ContentType          string             `json:"contentType"`
	Rating               int                `json:"rating"`
	StartTime            time.Time          `json:"startTime"`
	CompletedTime        time.Time          `json:"completedTime"`
	State                model.ArtworkState `json:"state"`
}

type Size struct {
	Width  float64 `json:"width" bson:"width"`
	Height float64 `json:"height" bson:"height"`
	Unit   string  `json:"unit" bson:"unit"`
}

func (c CreatedArtwork) ToDomainArtworkCreator() model.ArtworkCreator {
	return model.ArtworkCreator{
		CommissionID:      c.CommissionID,
		OpenCommissionID:  c.OpenCommissionID,
		ArtistID:          c.ArtistID,
		ArtistName:        c.ArtistName,
		ArtistProfilePath: c.ArtistProfilePath,
		DayUsed:           c.DayUsed,
		IsR18:             c.IsR18,
		Anonymous:         c.Anonymous,
		Path:              c.Path,
		Rating:            c.Rating,
		StartTime:         c.StartTime,
		CompletedTime:     c.CompletedTime,
	}
}
