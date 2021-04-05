package msg

import (
	"pixstall-search/domain/artist/model"
	"time"
)

type CreatedArtist struct {
	User           `json:",inline"`
	ArtistID       string            `json:"artistId"`
	ArtistIntro    model.ArtistIntro `json:"artistIntro"`
	PaymentMethods []string          `json:"paymentMethods"`
}

type User struct {
	UserID          string          `json:"userId"`
	UserName        string          `json:"userName"`
	ProfilePath     *string         `json:"profilePath"`
	Email           string          `json:"email"`
	Birthday        string          `json:"birthday"`
	Gender          string          `json:"gender"`
	State           model.UserState `json:"state"`
	RegTime         time.Time       `json:"regTime"`
	LastUpdatedTime time.Time       `json:"lastUpdatedTime"`
}

func (c CreatedArtist) ToDomainArtistCreator() model.ArtistCreator {
	return model.ArtistCreator{
		User: model.User{
			UserID:          c.UserID,
			UserName:        c.UserName,
			ProfilePath:     c.ProfilePath,
			State:           c.State,
			LastUpdatedTime: c.LastUpdatedTime,
		},
		ArtistID:       c.ArtistID,
		ArtistIntro:    c.ArtistIntro,
		PaymentMethods: c.PaymentMethods,
	}
}
