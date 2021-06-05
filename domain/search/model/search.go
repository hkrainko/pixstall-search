package model


type Search struct {
	SearchType SearchType

}




func NewSearchType(s string) *SearchType {
	var searchType SearchType
	switch s {
	case "all":
		searchType = AllSearchType
		return &searchType
	case "open-commissions":
		searchType = OpenCommissionsSearchType
		return &searchType
	case "artists":
		searchType = ArtistsSearchType
		return &searchType
	case "artworks":
		searchType = ArtworksSearchType
		return &searchType
	default:
		return nil
	}
}

type SearchType string

const (
	AllSearchType             SearchType = "all"
	OpenCommissionsSearchType SearchType = "open-commissions"
	ArtistsSearchType         SearchType = "artists"
	ArtworksSearchType        SearchType = "artworks"
)
