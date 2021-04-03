package msg

import "pixstall-search/domain/artist/model"

type UpdatedArtist struct {
	ArtistID string

	UserName    *string
	ProfilePath *string
	Email       *string
	Birthday    *string
	Gender      *string
	State       *string
	RegTime     *string

	ArtistIntro    *model.ArtistIntro
	ArtistBoard    *model.ArtistBoard
	PaymentMethods *[]string

	CommissionDetails *model.CommissionDetails
}

func (u UpdatedArtist) ToDomainArtistUpdater() model.ArtistUpdater {
	return model.ArtistUpdater{
		ArtistID:          u.ArtistID,
		UserName:          u.UserName,
		ProfilePath:       u.ProfilePath,
		State:             u.State,
		RegTime:           u.RegTime,
		ArtistIntro:       u.ArtistIntro,
		ArtistBoard:       u.ArtistBoard,
		PaymentMethods:    u.PaymentMethods,
		CommissionDetails: u.CommissionDetails,
	}
}
