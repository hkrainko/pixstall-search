package elastic_search

import (
	"context"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	elastic_search "pixstall-search/app/document/repo/elastic-search"
	"pixstall-search/app/search/repo/elastic-search/req"
	resp2 "pixstall-search/app/search/repo/elastic-search/resp"
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/artwork/model"
	error2 "pixstall-search/domain/error"
	model3 "pixstall-search/domain/open-commission/model"
	"pixstall-search/domain/search"
)

type elasticSearchSearchRepo struct {
	host   elastic_search.ElasticSearchHost
	client *resty.Client
}

func NewElasticSearchSearchRepo(host elastic_search.ElasticSearchHost) search.Repo {
	return &elasticSearchSearchRepo{
		host:   host,
		client: resty.New(),
	}
}

func (e elasticSearchSearchRepo) GetSuggestions(ctx context.Context, query string) (*[]string, error) {
	panic("implement me")
}

func (e elasticSearchSearchRepo) SearchArtists(ctx context.Context, query string, filter model.ArtistFilter, sorter model.ArtistSorter) (*model.GetArtistsResult, error) {
	var resp resp2.SearchArtistsResponse
	r, err := e.client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", e.host.BearToken()).
		SetBody(req.NewSearchArtistRequest(query, filter, sorter)).
		SetResult(&resp).
		Post(e.host.ApiPath + "/artists-search-engine/search")

	log.Println(r)
	log.Println(err)
	if err := checkIfError(r, err); err != nil {
		return nil, err
	}
	result := resp.ToDomainResult()
	return &result, nil
}

func (e elasticSearchSearchRepo) SearchArtworks(ctx context.Context, query string, filter model2.ArtworkFilter, sorter model2.ArtworkSorter) (*[]model2.Artwork, error) {
	panic("implement me")
}

func (e elasticSearchSearchRepo) SearchOpenCommissions(ctx context.Context, query string, filter model3.OpenCommissionFilter, sorter model3.OpenCommissionSorter) (*[]model3.OpenCommission, error) {
	panic("implement me")
}

func checkIfError(resp *resty.Response, err error) error {
	if err != nil {
		log.Println(err)
		return error2.UnknownError
	}
	log.Println(resp)
	if resp.StatusCode() != 200 {
		switch resp.StatusCode() {
		case http.StatusNotFound:
			return error2.NotFoundError
		default:
			return error2.UnknownError
		}
	}
	return nil
}
