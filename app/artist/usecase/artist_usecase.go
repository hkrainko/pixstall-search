package usecase

import (
	"context"
	"pixstall-search/domain/artist"
	"pixstall-search/domain/artist/model"
)

type artistUseCase struct {

}

func NewArtistUseCase() artist.UseCase {
	return &artistUseCase{}
}

func (a artistUseCase) CreateArtist(ctx context.Context, creator model.ArtistCreator) (*string, error) {
	panic("implement me")
}

func (a artistUseCase) UpdateArtist(ctx context.Context, updater model.ArtistUpdater) (*string, error) {
	panic("implement me")
}