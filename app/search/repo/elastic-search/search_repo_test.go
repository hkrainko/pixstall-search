package elastic_search

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"os"
	elastic_search "pixstall-search/app/document/repo/elastic-search"
	"pixstall-search/domain/artist/model"
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
	filter := model.ArtistFilter{
		State:                  nil,
		RegTime:                nil,
		PaymentMethods:         nil,
		YearOfDrawing:          nil,
		CommissionRequestCount: nil,
		CommissionAcceptCount:  nil,
		CommissionSuccessCount: nil,
		AvgRatings:             nil,
		LastRequestTime:        nil,
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
	artists, err := repo.SearchArtists(ctx, "artist", filter, sorter)
	assert.NoError(t, err)
	assert.NotNil(t, artists)
	if artists != nil {
		assert.Greater(t, len(*artists), 0)
	}
}
