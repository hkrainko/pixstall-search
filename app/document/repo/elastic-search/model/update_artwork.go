package model

import (
	"pixstall-search/domain/artwork/model"
	"time"
)

type UpdateArtworkRequest struct {
	ID                string  `json:"id"`
	ArtistName        *string `json:"artist_name,omitempty"`
	ArtistProfilePath *string `json:"artist_profile_path,omitempty"`

	Title       *string `json:"title,omitempty"`
	TextContent *string `json:"text_content,omitempty"`
	Views       *int    `json:"views,omitempty"`
	FavorCount  *int    `json:"favor_count,omitempty"`

	LastUpdateTime *time.Time          `json:"last_update_time,omitempty"`
	State          *model.ArtworkState `json:"state,omitempty"`
}

func NewUpdateArtworkRequest(u model.ArtworkUpdater) UpdateArtworkRequest {
	return UpdateArtworkRequest{
		ID:                u.ID,
		ArtistName:        u.ArtistName,
		ArtistProfilePath: u.ArtistProfilePath,
		Title:             u.Title,
		TextContent:       u.TextContent,
		Views:             u.Views,
		FavorCount:        u.FavorCount,
		LastUpdateTime:    u.LastUpdateTime,
		State:             u.State,
	}
}

type UpdateArtworkResponse struct {
	ID     string   `json:"id"`
	Errors []string `json:"errors"`
}
