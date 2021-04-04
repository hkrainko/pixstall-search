package usecase

import (
	"context"
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/artwork/model"
	model3 "pixstall-search/domain/open-commission/model"
	"pixstall-search/domain/search"
)

type searchUseCase struct {

}

func NewSearchUseCase() search.UseCase {
	return &searchUseCase{

	}
}

func (s searchUseCase) GetSuggestions(ctx context.Context, query string) (*[]string, error) {
	panic("implement me")
}

func (s searchUseCase) SearchArtists(ctx context.Context, query string, filter model.ArtistFilter, sorter model.ArtistSorter) (*[]model.Artist, error) {
	panic("implement me")
}

func (s searchUseCase) SearchArtworks(ctx context.Context, query string, filter model2.ArtworkFilter, sorter model2.ArtworkSorter) (*[]model2.Artwork, error) {
	panic("implement me")
}

func (s searchUseCase) SearchOpenCommissions(ctx context.Context, query string, filter model3.OpenCommissionFilter, sorter model3.OpenCommissionSorter) (*[]model3.OpenCommission, error) {
	panic("implement me")
}