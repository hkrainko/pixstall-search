package elastic_search

import (
	"context"
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/artwork/model"
	model3 "pixstall-search/domain/open-commission/model"
	"pixstall-search/domain/search"
)

type elasticSearchSearchRepo struct {

}

func NewElasticSearchSearchRepo() search.Repo {
	return &elasticSearchSearchRepo{

	}
}

func (e elasticSearchSearchRepo) GetSuggestions(ctx context.Context, query string) (*[]string, error) {
	panic("implement me")
}

func (e elasticSearchSearchRepo) SearchArtists(ctx context.Context, query string, filter model.ArtistFilter, sorter model.ArtistSorter) (*[]model.Artist, error) {
	panic("implement me")
}

func (e elasticSearchSearchRepo) SearchArtworks(ctx context.Context, query string, filter model2.ArtworkFilter, sorter model2.ArtworkSorter) (*[]model2.Artwork, error) {
	panic("implement me")
}

func (e elasticSearchSearchRepo) SearchOpenCommissions(ctx context.Context, query string, filter model3.OpenCommissionFilter, sorter model3.OpenCommissionSorter) (*[]model3.OpenCommission, error) {
	panic("implement me")
}