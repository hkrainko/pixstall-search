package elastic_search

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/artwork/model"
	"pixstall-search/domain/document"
	model3 "pixstall-search/domain/open-commission/model"
)

type elasticSearchDocumentRepo struct {

}

func NewElasticSearchDocumentRepo() document.Repo {
	return &elasticSearchDocumentRepo{

	}
}

func (e elasticSearchDocumentRepo) AddArtist(ctx context.Context, creator model.ArtistCreator) (*string, error) {
	var jsonStr []byte
	req, err := http.NewRequest("POST", "", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("", "")
	req.Header.Set("", "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body);
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