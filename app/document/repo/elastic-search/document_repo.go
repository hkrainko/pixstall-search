package elastic_search

import (
	"context"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
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
	client *resty.Client
}

func NewElasticSearchDocumentRepo(host ElasticSearchHost) document.Repo {
	return &elasticSearchDocumentRepo{
		host: host,
		client: resty.New(),
	}
}

func (e elasticSearchDocumentRepo) AddArtist(ctx context.Context, creator model.ArtistCreator) (*string, error) {
	var resp []model4.ResponseResult
	r, err := e.client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", e.host.BearToken()).
		SetBody(model4.NewAddArtistRequest(creator)).
		SetResult(&resp).
		Post(e.host.ApiPath + "/artists-search-engine/documents")
	if err := checkIfError(r, err); err != nil {
		return nil, err
	}
	if len(resp[0].Errors) > 0 {
		log.Printf("%v", resp[0].Errors)
		return nil, error2.UnknownError
	}
	return &resp[0].ID, nil
}

func (e elasticSearchDocumentRepo) UpdateArtist(ctx context.Context, updater model.ArtistUpdater) (*string, error) {
	var result []model4.ResponseResult
	r, err := e.client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", e.host.BearToken()).
		SetBody(model4.NewUpdateArtistRequest(updater)).
		SetResult(&result).
		Patch(e.host.ApiPath + "/artists-search-engine/documents")
	if err := checkIfError(r, err); err != nil {
		return nil, err
	}
	if len(result[0].Errors) > 0 {
		log.Printf("%v", result[0].Errors)
		return nil, error2.UnknownError
	}
	return &result[0].ID, nil
}

func (e elasticSearchDocumentRepo) AddArtwork(ctx context.Context, creator model2.ArtworkCreator) (*string, error) {
	var resp []model4.ResponseResult
	r, err := e.client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", e.host.BearToken()).
		SetBody(model4.NewAddArtworkRequest(creator)).
		SetResult(&resp).
		Post(e.host.ApiPath + "/artworks-search-engine/documents")
	log.Println(r)
	if err := checkIfError(r, err); err != nil {
		return nil, err
	}
	if len(resp[0].Errors) > 0 {
		log.Printf("%v", resp[0].Errors)
		return nil, error2.UnknownError
	}
	return &resp[0].ID, nil
}

func (e elasticSearchDocumentRepo) UpdateArtwork(ctx context.Context, updater model2.ArtworkUpdater) (*string, error) {
	var result []model4.ResponseResult
	r, err := e.client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", e.host.BearToken()).
		SetBody(model4.NewUpdateArtworkRequest(updater)).
		SetResult(&result).
		Patch(e.host.ApiPath + "/artworks-search-engine/documents")
	if err := checkIfError(r, err); err != nil {
		return nil, err
	}
	if len(result[0].Errors) > 0 {
		log.Printf("%v", result[0].Errors)
		return nil, error2.UnknownError
	}
	return &result[0].ID, nil
}

func (e elasticSearchDocumentRepo) AddOpenCommission(ctx context.Context, creator model3.OpenCommissionCreator) (*string, error) {
	var resp []model4.ResponseResult
	r, err := e.client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", e.host.BearToken()).
		SetBody(model4.NewAddOpenCommissionRequest(creator)).
		SetResult(&resp).
		Post(e.host.ApiPath + "/open-commissions-search-engine/documents")
	log.Println(r)
	if err := checkIfError(r, err); err != nil {
		return nil, err
	}
	if len(resp[0].Errors) > 0 {
		log.Printf("%v", resp[0].Errors)
		return nil, error2.UnknownError
	}
	return &resp[0].ID, nil
}

func (e elasticSearchDocumentRepo) UpdateOpenCommission(ctx context.Context, updater model3.OpenCommissionUpdater) (*string, error) {
	var result []model4.ResponseResult
	r, err := e.client.
		R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", e.host.BearToken()).
		SetBody(model4.NewUpdateOpenCommissionRequest(updater)).
		SetResult(&result).
		Patch(e.host.ApiPath + "/open-commissions-search-engine/documents")
	if err := checkIfError(r, err); err != nil {
		return nil, err
	}
	if len(result[0].Errors) > 0 {
		log.Printf("%v", result[0].Errors)
		return nil, error2.UnknownError
	}
	return &result[0].ID, nil
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