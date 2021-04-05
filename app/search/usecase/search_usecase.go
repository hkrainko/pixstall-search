package usecase

import (
	"context"
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/artwork/model"
	model3 "pixstall-search/domain/open-commission/model"
	"pixstall-search/domain/search"
)

type searchUseCase struct {
	repo search.Repo
}

func NewSearchUseCase(repo search.Repo) search.UseCase {
	return &searchUseCase{
		repo: repo,
	}
}

func (s searchUseCase) GetSuggestions(ctx context.Context, query string) (*[]string, error) {
	panic("implement me")
}

func (s searchUseCase) SearchArtists(ctx context.Context, query string, filter model.ArtistFilter, sorter model.ArtistSorter) (*model.GetArtistsResult, error) {
	return s.repo.SearchArtists(ctx, query, filter, sorter)
}

func (s searchUseCase) SearchArtworks(ctx context.Context, query string, filter model2.ArtworkFilter, sorter model2.ArtworkSorter) (*model2.GetArtworksResult, error) {
	return s.repo.SearchArtworks(ctx, query, filter, sorter)
}

func (s searchUseCase) SearchOpenCommissions(ctx context.Context, query string, filter model3.OpenCommissionFilter, sorter model3.OpenCommissionSorter) (*model3.GetOpenCommissionsResult, error) {
	return s.repo.SearchOpenCommissions(ctx, query, filter, sorter)
}
