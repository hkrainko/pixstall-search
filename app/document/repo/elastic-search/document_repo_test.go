package elastic_search

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"pixstall-search/domain/artist/model"
	"pixstall-search/domain/document"
	"testing"
	"time"
)

var repo document.Repo
var ctx context.Context

const apiPath = "http://localhost:3002/api/as/v1/engines"
const key = "search-21ovx6yqaffz92sxw8qpu23p"
const token = "private-jxcq12x5tuko6rkbh28gkbdk"

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("Before all tests")
	ctx = context.Background()
	repo = NewElasticSearchDocumentRepo(ElasticSearchHost{
		ApiPath: apiPath,
		Key: key,
		token: token,
	})
}

func teardown() {
	fmt.Println("After all tests")
}

func TestAddArtist(t *testing.T) {
	artistCreator := model.ArtistCreator{
		User:           model.User{
			UserID:          "userID",
			UserName:        "userName",
			ProfilePath:     "profilePath",
			State:           model.UserStateActive,
			LastUpdatedTime: time.Now(),
		},
		ArtistID:       "artistID",
		ArtistIntro:    model.ArtistIntro{
			YearOfDrawing: 5,
			ArtTypes:      []string{"A", "B"},
		},
		PaymentMethods: []string{"paypal"},
	}
	id, err := repo.AddArtist(ctx, artistCreator)
	assert.NoError(t, err)
	assert.Equal(t, "artistID", *id)
}

func TestUpdateArtist(t *testing.T) {

}