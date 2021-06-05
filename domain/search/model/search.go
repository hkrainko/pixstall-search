package model

type SearchType string

const (
	AllSearchType             SearchType = "all"
	OpenCommissionsSearchType SearchType = "open-commissions"
	ArtistsSearchType         SearchType = "artists"
	ArtworksSearchType        SearchType = "artworks"
)
