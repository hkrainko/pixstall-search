package usecase

import (
	"context"
	"pixstall-search/domain/artwork"
	"pixstall-search/domain/artwork/model"
)

type artworkUseCase struct {

}

func NewArtworkUseCase() artwork.UseCase {
	return &artworkUseCase{}
}

func (a artworkUseCase) CreateArtwork(ctx context.Context, creator model.ArtworkCreator) (*string, error) {
	panic("implement me")
}

func (a artworkUseCase) UpdateArtwork(ctx context.Context, updater model.ArtworkUpdater) (*string, error) {
	panic("implement me")
}