package elastic_search

import (
	"context"
	"github.com/go-resty/resty/v2"
	"log"
	model4 "pixstall-search/app/document/repo/elastic-search/model"
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/artwork/model"
	"pixstall-search/domain/document"
	error2 "pixstall-search/domain/error"
	model3 "pixstall-search/domain/open-commission/model"
)

type ElasticSearchHost struct {
	ApiPath string
	Key string
	token string
}

func (e ElasticSearchHost) BearToken() string {
	return "Bearer " + e.token
}

type elasticSearchDocumentRepo struct {
	host ElasticSearchHost
}

func NewElasticSearchDocumentRepo(host ElasticSearchHost) document.Repo {
	return &elasticSearchDocumentRepo{
		host: host,
	}
}

func (e elasticSearchDocumentRepo) AddArtist(ctx context.Context, creator model.ArtistCreator) (*string, error) {
	client := resty.New()

	var resp []model4.AddArtistResponse

	_, err := client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", e.host.BearToken()).
		SetBody(model4.NewAddArtistRequestFromArtistCreator(creator)).
		SetResult(&resp).
		Post(e.host.ApiPath + "/artist-search-engine/documents")
	if err != nil {
		log.Println(err)
		return nil, error2.UnknownError
	}
	if len(resp[0].Errors) > 0 {
		log.Printf("%v", resp[0].Errors)
		return nil, error2.UnknownError
	}
	return &resp[0].ID, nil
}

func (e elasticSearchDocumentRepo) UpdateArtist(ctx context.Context, updater model.ArtistUpdater) (*string, error) {
	panic("implement me")
}

func (e elasticSearchDocumentRepo) AddArtwork(ctx context.Context, creator model2.ArtworkCreator) (*string, error) {
	panic("implement me")
}

func (e elasticSearchDocumentRepo) UpdateArtwork(ctx context.Context, updater model2.ArtworkUpdater) (*string, error) {
	panic("implement me")
}

func (e elasticSearchDocumentRepo) AddOpenCommission(ctx context.Context, creator model3.OpenCommissionCreator) (*string, error) {
	panic("implement me")
}

func (e elasticSearchDocumentRepo) UpdateOpenCommission(ctx context.Context, updater model3.OpenCommissionUpdater) (*string, error) {
	panic("implement me")
}