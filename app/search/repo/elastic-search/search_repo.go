package elastic_search

import (
	"context"
	"github.com/go-resty/resty/v2"
	"log"
	elastic_search "pixstall-search/app/document/repo/elastic-search"
	model4 "pixstall-search/app/search/repo/elastic-search/model"
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/artwork/model"
	error2 "pixstall-search/domain/error"
	model3 "pixstall-search/domain/open-commission/model"
	"pixstall-search/domain/search"
)

type elasticSearchSearchRepo struct {
	host elastic_search.ElasticSearchHost
	client *resty.Client
}

func NewElasticSearchSearchRepo(host elastic_search.ElasticSearchHost) search.Repo {
	return &elasticSearchSearchRepo{
		host: host,
		client: resty.New(),
	}
}

func (e elasticSearchSearchRepo) GetSuggestions(ctx context.Context, query string) (*[]string, error) {
	panic("implement me")
}

func (e elasticSearchSearchRepo) SearchArtists(ctx context.Context, query string, filter model.ArtistFilter, sorter model.ArtistSorter) (*[]model.Artist, error) {
	var resp model4.SearchArtistsResponse
	r, err := e.client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", e.host.BearToken()).
		SetBody(model4.NewSearchArtistRequest(query, filter, sorter)).
		SetResult(&resp).
		Post(e.host.ApiPath + "/artists-search-engine/search")
	if err := checkIfError(r, err); err != nil {
		return nil, err
	}
	if len(resp[0].Errors) > 0 {
		log.Printf("%v", resp[0].Errors)
		return nil, error2.UnknownError
	}
	return &resp[0].ID, nil
}

func (e elasticSearchSearchRepo) SearchArtworks(ctx context.Context, query string, filter model2.ArtworkFilter, sorter model2.ArtworkSorter) (*[]model2.Artwork, error) {
	panic("implement me")
}

func (e elasticSearchSearchRepo) SearchOpenCommissions(ctx context.Context, query string, filter model3.OpenCommissionFilter, sorter model3.OpenCommissionSorter) (*[]model3.OpenCommission, error) {
	panic("implement me")
}