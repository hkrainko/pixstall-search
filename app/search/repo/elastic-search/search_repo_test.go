package elastic_search

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"os"
	elastic_search "pixstall-search/app/document/repo/elastic-search"
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/model"
	"pixstall-search/domain/search"
	"testing"
)

var repo search.Repo
var ctx context.Context
var faker *gofakeit.Faker

const apiPath = "http://localhost:3002/api/as/v1/engines"
const key = "search-21ovx6yqaffz92sxw8qpu23p"

func TestMain(m *testing.M) {
	faker = gofakeit.New(0)
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("Before eacg tests")
	ctx = context.Background()
	repo = NewElasticSearchSearchRepo(elastic_search.ElasticSearchHost{
		ApiPath: apiPath,
		Key:     key,
	})
}

func teardown() {
	fmt.Println("After all tests")
}

func TestSearchArtists(t *testing.T) {
	state := []model.UserState{model.UserStateActive}
	filter := model.ArtistFilter{
		State:                  &state,
		RegTime:                nil,
		PaymentMethods:         nil,
		YearOfDrawing:          nil,
		CommissionRequestCount: nil,
		CommissionAcceptCount:  nil,
		CommissionSuccessCount: nil,
		AvgRatings:             nil,
		LastRequestTime:        nil,
		PageFilter: model2.PageFilter{
			Current: 1,
			Size:    50,
		},
	}
	sorter := model.ArtistSorter{
		ArtistID:               nil,
		UserName:               nil,
		RegTime:                nil,
		YearOfDrawing:          nil,
		CommissionRequestCount: nil,
		CommissionAcceptCount:  nil,
		CommissionSuccessCount: nil,
		AvgRatings:             nil,
		LastRequestTime:        nil,
	}
	getArtistsResult, err := repo.SearchArtists(ctx, "artist", filter, sorter)
	assert.NoError(t, err)
	assert.NotNil(t, getArtistsResult)
	if getArtistsResult != nil {
		assert.Greater(t, len(getArtistsResult.Artists), 0)
	}
}
