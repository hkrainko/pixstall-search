package usecase

import (
	"context"
	"pixstall-search/domain/artwork"
	"pixstall-search/domain/artwork/model"
	"pixstall-search/domain/document"
)

type artworkUseCase struct {
	docRepo document.Repo
}

func NewArtworkUseCase(docRepo document.Repo) artwork.UseCase {
	return &artworkUseCase{
		docRepo: docRepo,
	}
}

func (a artworkUseCase) CreateArtwork(ctx context.Context, creator model.ArtworkCreator) (*string, error) {
	return a.docRepo.AddArtwork(ctx, creator)
}

func (a artworkUseCase) UpdateArtwork(ctx context.Context, updater model.ArtworkUpdater) (*string, error) {
	return a.docRepo.UpdateArtwork(ctx, updater)
}
