package document

import (
	"context"
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/artwork/model"
	model3 "pixstall-search/domain/open-commission/model"
)

type Repo interface {
	AddArtist(ctx context.Context, creator model.ArtistCreator) (*string, error)
	UpdateArtist(ctx context.Context, updater model.ArtistUpdater) (*string, error)

	AddArtwork(ctx context.Context, creator model2.ArtworkCreator) (*string, error)
	UpdateArtwork(ctx context.Context, updater model2.ArtworkUpdater) (*string, error)

	AddOpenCommission(ctx context.Context, creator model3.OpenCommissionCreator) (*string, error)
	UpdateOpenCommission(ctx context.Context, updater model3.OpenCommissionUpdater) (*string, error)

}
