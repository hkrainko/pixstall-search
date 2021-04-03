package model

import (
	"pixstall-search/domain/artist/model"
	"time"
)

type UpdateArtistRequest struct {
	ID string

	UserID          *string          `json:"user_id"`
	UserName        *string          `json:"user_name"`
	ProfilePath     *string          `json:"profile_path"`
	State           *model.UserState `json:"state"`
	RegTime         *time.Time       `json:"reg_time"`
	LastUpdatedTime *time.Time       `json:"last_updated_time"`

	ArtistIntroYearOfDrawing *int
	ArtistIntroArtTypes      *[]string
	PaymentMethods           *[]string

	CommissionDetails *CommissionDetails
}
