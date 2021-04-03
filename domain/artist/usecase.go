package artist

import (
	"context"
	"pixstall-search/domain/artist/model"
)

type UseCase interface {
	CreateArtist(ctx context.Context, creator model.ArtistCreator) (*string, error)
	UpdateArtist(ctx context.Context, updater model.ArtistUpdater) (*string, error)
}