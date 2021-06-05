package elastic_search

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"os"
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/artwork/model"
	"pixstall-search/domain/document"
	model3 "pixstall-search/domain/open-commission/model"
	"strconv"
	"testing"
	"time"
)

var repo document.Repo
var ctx context.Context
var faker *gofakeit.Faker

const apiPath = "https://kai-dev-deployment.ent.ap-southeast-1.aws.found.io/api/as/v1/engines"
const key = "private-1xkbqfccp6xkrfo5t8j3vpia"

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
	repo = NewElasticSearchDocumentRepo(&ElasticSearchHost{
		ApiPath:    apiPath,
		PrivateKey: key,
	})
}

func teardown() {
	fmt.Println("After all tests")
}

func TestAddArtist(t *testing.T) {
	profilePath := "profilePath"
	artistCreator := model.ArtistCreator{
		User: model.User{
			UserID:          "artistID",
			UserName:        "userName",
			ProfilePath:     &profilePath,
			State:           model.UserStateActive,
			RegTime:         time.Now(),
			LastUpdatedTime: time.Now(),
		},
		ArtistID: "artistID",
		ArtistIntro: model.ArtistIntro{
			YearOfDrawing: 5,
			ArtTypes:      []string{"A", "B"},
		},
		PaymentMethods: []string{"paypal"},
	}
	id, err := repo.AddArtist(ctx, artistCreator)
	assert.NoError(t, err)
	assert.Equal(t, "artistID", *id)
}

func TestAddArtist_multiple(t *testing.T) {
	var err error
	for i := range [1000]int{} {
		profilePath := faker.ImageURL(200, 100)
		artistCreator := model.ArtistCreator{
			User: model.User{
				UserID:          "artistID" + strconv.Itoa(i),
				UserName:        faker.Name(),
				ProfilePath:     &profilePath,
				State:           model.UserStateActive,
				RegTime:         faker.Date(),
				LastUpdatedTime: faker.Date(),
			},
			ArtistID: "artistID" + strconv.Itoa(i),
			ArtistIntro: model.ArtistIntro{
				YearOfDrawing: faker.Number(0, 20),
				ArtTypes:      []string{faker.Word(), faker.Word(), faker.Word()},
			},
			PaymentMethods: []string{faker.Word()},
		}
		_, err = repo.AddArtist(ctx, artistCreator)
	}
	assert.NoError(t, err)
}

func TestUpdateArtist(t *testing.T) {
	userName := "userName_update"
	profilePath := "profilePath_update"
	state := model.UserStateTerminated
	now := time.Now()
	avgRatings := 3.2
	yearOfDrawing := 999
	artTypes := []string{"Type1_update", "Type2_update"}
	commissionRequestCount := 9
	commissionAcceptCount := 9
	commissionSuccessCount := 9

	desc := "Desc_update"
	artistUpdater := model.ArtistUpdater{
		ArtistID:               "artistID",
		UserName:               &userName,
		ProfilePath:            &profilePath,
		State:                  &state,
		YearOfDrawing:          &yearOfDrawing,
		ArtTypes:               &artTypes,
		Desc:                   &desc,
		PaymentMethods:         &[]string{"payment_update", "payment_update__"},
		CommissionRequestCount: &commissionRequestCount,
		CommissionAcceptCount:  &commissionAcceptCount,
		CommissionSuccessCount: &commissionSuccessCount,
		AvgRatings:             &avgRatings,
		LastRequestTime:        &now,
		LastUpdatedTime:        &now,
	}
	id, err := repo.UpdateArtist(ctx, artistUpdater)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	if id != nil {
		assert.Equal(t, "artistID", *id)
	}
}

func TestAddArtwork(t *testing.T) {
	artistProfilePath := faker.ImageURL(100, 100)
	creator := model2.ArtworkCreator{
		ID:                "artworkID",
		CommissionID:      faker.UUID(),
		OpenCommissionID:  faker.UUID(),
		ArtistID:          faker.UUID(),
		ArtistName:        faker.Name(),
		ArtistProfilePath: &artistProfilePath,
		DayUsed:           time.Duration(faker.Day()),
		IsR18:             faker.Bool(),
		Anonymous:         faker.Bool(),
		Path:              faker.ImageURL(1000, 1000),
		Rating:            faker.RandomInt([]int{1, 2, 3, 4, 5}),
		CreateTime:        faker.Date(),
		StartTime:         faker.Date(),
		CompletedTime:     faker.Date(),
		State:             model2.ArtworkStateActive,
	}
	id, err := repo.AddArtwork(ctx, creator)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	if id != nil {
		assert.Equal(t, "artworkID", *id)
	}
}

func TestAddArtwork_multiple(t *testing.T) {
	var err error
	for i := range [100]int{} {
		artistProfilePath := faker.ImageURL(100, 100)
		creator := model2.ArtworkCreator{
			ID:                "artworkID" + strconv.Itoa(i),
			CommissionID:      faker.UUID(),
			OpenCommissionID:  faker.UUID(),
			ArtistID:          faker.UUID(),
			ArtistName:        faker.Name(),
			ArtistProfilePath: &artistProfilePath,
			DayUsed:           time.Duration(faker.Day()),
			IsR18:             faker.Bool(),
			Anonymous:         faker.Bool(),
			Path:              faker.ImageURL(1000, 1000),
			Rating:            faker.RandomInt([]int{1, 2, 3, 4, 5}),
			CreateTime:        faker.Date(),
			StartTime:         faker.Date(),
			CompletedTime:     faker.Date(),
			State:             model2.ArtworkStateActive,
		}
		_, err = repo.AddArtwork(ctx, creator)
	}
	assert.NoError(t, err)
}

func TestUpdateArtwork(t *testing.T) {
	artistName := faker.Name()
	artistProfilePath := faker.ImageURL(100, 100)
	title := faker.Sentence(10)
	textContent := faker.LoremIpsumParagraph(3, 5, 10, "\n")
	views := faker.Number(100, 100000)
	favorCount := faker.Number(100, 100000)
	lastUpdatedTime := faker.Date()
	state := model2.ArtworkStateDeleted
	updater := model2.ArtworkUpdater{
		ID:                "artworkID",
		ArtistName:        &artistName,
		ArtistProfilePath: &artistProfilePath,
		Title:             &title,
		TextContent:       &textContent,
		Views:             &views,
		FavorCount:        &favorCount,
		LastUpdatedTime:   &lastUpdatedTime,
		State:             &state,
	}
	id, err := repo.UpdateArtwork(ctx, updater)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	if id != nil {
		assert.Equal(t, "artworkID", *id)
	}
}

func TestAddOpenCommission(t *testing.T) {

	creator := model3.OpenCommissionCreator{
		ID:       "openCommissionID",
		ArtistID: faker.UUID(),
		Title:    faker.Sentence(10),
		Desc:     faker.Paragraph(2, 3, 4, "\n"),
		Price: model3.Price{
			Amount:   faker.Price(100, 1000),
			Currency: model3.CurrencyHKD,
		},
		DayNeed: model3.DayNeed{
			From: faker.Number(10, 20),
			To:   faker.Number(20, 30),
		},
		IsR18:          faker.Bool(),
		AllowBePrivate: faker.Bool(),
		AllowAnonymous: faker.Bool(),
		SampleImagePaths: []string{
			faker.ImageURL(100, 100),
			faker.ImageURL(100, 100),
			faker.ImageURL(100, 100),
		},
		State:           model3.OpenCommissionStateActive,
		CreateTime:      faker.Date(),
		LastUpdatedTime: faker.Date(),
	}
	id, err := repo.AddOpenCommission(ctx, creator)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	if id != nil {
		assert.Equal(t, "openCommissionID", *id)
	}
}

func TestAddOpenCommissions_multiple(t *testing.T) {
	var err error
	for i := range [100]int{} {
		creator := model3.OpenCommissionCreator{
			ID:       "openCommissionID" + strconv.Itoa(i),
			ArtistID: faker.UUID(),
			Title:    faker.Sentence(10),
			Desc:     faker.Paragraph(2, 3, 4, "\n"),
			Price: model3.Price{
				Amount:   faker.Price(100, 1000),
				Currency: model3.CurrencyHKD,
			},
			DayNeed: model3.DayNeed{
				From: faker.Number(10, 20),
				To:   faker.Number(20, 30),
			},
			IsR18:          faker.Bool(),
			AllowBePrivate: faker.Bool(),
			AllowAnonymous: faker.Bool(),
			SampleImagePaths: []string{
				faker.ImageURL(100, 100),
				faker.ImageURL(100, 100),
				faker.ImageURL(100, 100),
			},
			State:           model3.OpenCommissionStateActive,
			CreateTime:      faker.Date(),
			LastUpdatedTime: faker.Date(),
		}
		_, err = repo.AddOpenCommission(ctx, creator)
	}
	assert.NoError(t, err)
}

func TestUpdateOpenCommission(t *testing.T) {
	title := faker.Sentence(20)
	desc := faker.Paragraph(2, 3, 10, "\n")
	isR18 := faker.Bool()
	allowBePrivate := faker.Bool()
	allowAnonymous := faker.Bool()
	state := model3.OpenCommissionStateRemoved

	updater := model3.OpenCommissionUpdater{
		ID:    "openCommissionID",
		Title: &title,
		Desc:  &desc,
		Price: &model3.Price{
			Amount:   faker.Price(100, 300),
			Currency: model3.CurrencyTWD,
		},
		DayNeed: &model3.DayNeed{
			From: faker.Number(10, 20),
			To:   faker.Number(20, 110),
		},
		SampleImagePaths: &[]string{
			faker.ImageURL(100, 100),
			faker.ImageURL(100, 100),
			faker.ImageURL(100, 100),
		},
		IsR18:           &isR18,
		AllowBePrivate:  &allowBePrivate,
		AllowAnonymous:  &allowAnonymous,
		State:           &state,
		LastUpdatedTime: faker.Date(),
	}
	id, err := repo.UpdateOpenCommission(ctx, updater)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	if id != nil {
		assert.Equal(t, "openCommissionID", *id)
	}
}
