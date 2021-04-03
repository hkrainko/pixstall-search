package artwork

import (
	"context"
	"pixstall-search/domain/artwork/model"
)

type UseCase interface {
	CreateArtwork(ctx context.Context, creator model.ArtworkCreator) (*string, error)
	UpdateArtwork(ctx context.Context, updater model.ArtworkUpdater) (*string, error)
}