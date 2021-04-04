package model

import "pixstall-search/domain/artist/model"

type SearchArtistsRequest struct {

	Filters map[string]string `json:"filters"`
	Sort map[string]string `json:"sort"`
}

type SearchArtistsResponse struct {

}

func NewSearchArtistRequest(query string, filter model.ArtistFilter, sorting model.ArtistSorter) SearchArtistsRequest {

	f := getFilterMap(filter)



	return SearchArtistsRequest{

		Filters: nil,
		Sort: nil,
	}
}

func getFilterMap(filter model.ArtistFilter) map[string]string {
	f := make(map[string]interface{})
	f["state"] = []model.UserState{model.UserStateActive}
	if filter.RegTimeFrom != nil {
		from :=
	}

}

type TimeRangeMap struct {
	*From
}