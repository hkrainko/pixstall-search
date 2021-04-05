package msg

import (
	"pixstall-search/domain/artwork/model"
	"time"
)

type UpdatedArtwork struct {
	ID                   string  `json:"id"`
	ArtistName           *string `json:"artistName"`
	ArtistProfilePath    *string `json:"artistProfilePath"`
	RequesterName        *string `json:"requesterName"`
	RequesterProfilePath *string `json:"requesterProfilePath"`

	Title       *string `json:"title"`
	TextContent *string `json:"textContent"`
	Views       *int    `json:"views"`
	FavorCount  *int    `json:"favorCount"`

	LastUpdatedTime *time.Time          `json:"lastUpdatedTime"`
	State           *model.ArtworkState `json:"state"`
}

func (u UpdatedArtwork) ToDomainArtworkUpdater() model.ArtworkUpdater {
	return model.ArtworkUpdater{
		ID:                u.ID,
		ArtistName:        u.ArtistName,
		ArtistProfilePath: u.ArtistProfilePath,
		Title:             u.Title,
		TextContent:       u.TextContent,
		Views:             u.Views,
		FavorCount:        u.FavorCount,
		LastUpdatedTime:   u.LastUpdatedTime,
		State:             u.State,
	}
}
