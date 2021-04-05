package usecase

import (
	"context"
	"pixstall-search/domain/artist"
	"pixstall-search/domain/artist/model"
	"pixstall-search/domain/document"
)

type artistUseCase struct {
	docRepo document.Repo
}

func NewArtistUseCase(docRepo document.Repo) artist.UseCase {
	return &artistUseCase{
		docRepo: docRepo,
	}
}

func (a artistUseCase) CreateArtist(ctx context.Context, creator model.ArtistCreator) (*string, error) {
	return a.docRepo.AddArtist(ctx, creator)
}

func (a artistUseCase) UpdateArtist(ctx context.Context, updater model.ArtistUpdater) (*string, error) {
	return a.docRepo.UpdateArtist(ctx, updater)
}