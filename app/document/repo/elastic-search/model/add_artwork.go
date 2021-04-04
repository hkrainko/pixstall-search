package model

import (
	"pixstall-search/domain/artwork/model"
	"time"
)

type AddArtworkRequest struct {
	ID               string `json:"id"`
	OpenCommissionID string `json:"open_commission_id"`

	ArtistID          string  `json:"artist_id"`
	ArtistName        string  `json:"artist_name"`
	ArtistProfilePath *string `json:"artist_profile_path"`

	DayUsed   time.Duration `json:"day_used"`
	IsR18     bool          `json:"is_r18"`
	Anonymous bool          `json:"anonymous"`

	Path   string `json:"path"`
	Rating int    `json:"rating"`

	Title       string `json:"title"`
	TextContent string `json:"text_content"`
	Views       int    `json:"views"`
	FavorCount  int    `json:"favor_count"`

	CreateTime     time.Time          `json:"create_time"`
	StartTime      time.Time          `json:"start_time"`
	CompletedTime  time.Time          `json:"completed_time"`
	LastUpdateTime time.Time          `json:"last_update_time"`
	State          model.ArtworkState `json:"state"`
}

func NewAddArtworkRequest(creator model.ArtworkCreator) AddArtworkRequest {
	return AddArtworkRequest{
		ID:                creator.ID,
		OpenCommissionID:  creator.OpenCommissionID,
		ArtistID:          creator.ArtistID,
		ArtistName:        creator.ArtistName,
		ArtistProfilePath: creator.ArtistProfilePath,
		DayUsed:           creator.DayUsed,
		IsR18:             creator.IsR18,
		Anonymous:         creator.Anonymous,
		Path:              creator.Path,
		Rating:            creator.Rating,
		Title:             "",
		TextContent:       "",
		Views:             0,
		FavorCount:        0,
		CreateTime:        creator.CreateTime,
		StartTime:         creator.StartTime,
		CompletedTime:     creator.CompletedTime,
		LastUpdateTime:    time.Now(),
		State:             creator.State,
	}
}

type AddArtworkResponse struct {
	ID     string   `json:"id"`
	Errors []string `json:"errors"`
}
