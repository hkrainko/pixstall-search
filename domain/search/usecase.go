package search

import (
	"context"
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/artwork/model"
	model3 "pixstall-search/domain/open-commission/model"
)

type UseCase interface {
	GetSuggestions(ctx context.Context, query string) (*[]string, error)
	SearchArtists(ctx context.Context, query string, filter model.ArtistFilter, sorter model.ArtistSorter) (*model.GetArtistsResult, error)
	SearchArtworks(ctx context.Context, query string, filter model2.ArtworkFilter, sorter model2.ArtworkSorter) (*model2.GetArtworksResult, error)
	SearchOpenCommissions(ctx context.Context, query string, filter model3.OpenCommissionFilter, sorter model3.OpenCommissionSorter) (*model3.GetOpenCommissionsResult, error)
}
