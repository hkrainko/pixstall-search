package elastic_search

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"os"
	elastic_search "pixstall-search/app/document/repo/elastic-search"
	"pixstall-search/domain/artist/model"
	model3 "pixstall-search/domain/artwork/model"
	model2 "pixstall-search/domain/model"
	model4 "pixstall-search/domain/open-commission/model"
	"pixstall-search/domain/search"
	"testing"
	"time"
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
	yearOfDrawingForm := 0
	yearOfDrawingTo := 100
	filter := model.ArtistFilter{
		State:                  &state,
		RegTime:                nil,
		PaymentMethods:         nil,
		YearOfDrawing:          &model2.IntRange{
			From: &yearOfDrawingForm,
			To:   &yearOfDrawingTo,
		},
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
	getArtistsResult, err := repo.SearchArtists(ctx, "resp", filter, sorter)
	assert.NoError(t, err)
	assert.NotNil(t, getArtistsResult)
	if getArtistsResult != nil {
		assert.Greater(t, len(getArtistsResult.Artists), 0)
	}
}

func TestSearchArtworks(t *testing.T) {
	state := []model3.ArtworkState{model3.ArtworkStateActive}
	dayUsed := time.Duration(5)
	filter := model3.ArtworkFilter{
		State:          &state,
		DayUsed:        &dayUsed,
		IsR18:          nil,
		Anonymous:      nil,
		Rating:         nil,
		Views:          nil,
		FavorCount:     nil,
		CreateTime:     nil,
		StartTime:      nil,
		CompletedTime:  nil,
		LastUpdateTime: nil,
		PageFilter:     model2.PageFilter{
			Current: 1,
			Size:    100,
		},
	}
	sorter := model3.ArtworkSorter{
		ArtistID:       nil,
		UserName:       nil,
		DayUsed:        nil,
		Rating:         nil,
		Views:          nil,
		FavorCount:     nil,
		CreateTime:     nil,
		StartTime:      nil,
		CompletedTime:  nil,
		LastUpdateTime: nil,
	}
	getArtworkResult, err := repo.SearchArtworks(ctx, "df", filter, sorter)
	assert.NoError(t, err)
	assert.NotNil(t, getArtworkResult)
	if getArtworkResult != nil {
		assert.Greater(t, len(getArtworkResult.Artworks), 0)
	}
}

func TestSearchOpenCommissions(t *testing.T) {
	state := []model4.OpenCommissionState{model4.OpenCommissionStateActive}
	priceAmountFrom := 10.00
	PriceAmountTo := 100.00
	priceCurrency := model4.CurrencyHKD
	dayNeedFromFrom := 5
	dayNeedFromTo := 20
	filter := model4.OpenCommissionFilter{
		State: &state,
		PriceAmount: &model2.FloatRange{
			From: &priceAmountFrom,
			To:   &PriceAmountTo,
		},
		PriceCurrency: &priceCurrency,
		DayNeedFrom: &model2.IntRange{
			From: &dayNeedFromFrom,
			To:   &dayNeedFromTo,
		},
		DayNeedTo: nil,
		IsR18:          nil,
		AllowBePrivate: nil,
		AllowAnonymous: nil,
		PageFilter: model2.PageFilter{
			Current: 1,
			Size:    50,
		},
	}
	sorter := model4.OpenCommissionSorter{
		ArtistID:        nil,
		Price:           nil,
		DayNeedFrom:     nil,
		DayNeedTo:       nil,
		CreateTime:      nil,
		LastUpdatedTime: nil,
	}
	getOpenCommResult, err := repo.SearchOpenCommissions(ctx, "open", filter, sorter)
	assert.NoError(t, err)
	assert.NotNil(t, getOpenCommResult)
	if getOpenCommResult != nil {
		assert.Greater(t, len(getOpenCommResult.OpenCommissions), 0)
	}
}
