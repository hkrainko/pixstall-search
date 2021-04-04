package model

import (
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/model"
)

type SearchArtistsRequest struct {
	Query   string               `json:"query"`
	Filters *SearchArtistsFilter `json:"filters,omitempty"`
	Sort    *SearchArtistsSorter `json:"sort,omitempty"`
}

type SearchArtistsResponse struct {
}

func NewSearchArtistRequest(query string, filter model.ArtistFilter, sorter model.ArtistSorter) SearchArtistsRequest {
	f := getFilter(filter)
	s := getSorter(sorter)
	return SearchArtistsRequest{
		Query:   query,
		Filters: &f,
		Sort:    &s,
	}
}

func getFilter(filter model.ArtistFilter) SearchArtistsFilter {
	return SearchArtistsFilter{
		State:                  filter.State,
		RegTime:                filter.RegTime,
		PaymentMethods:         filter.PaymentMethods,
		YearOfDrawing:          filter.YearOfDrawing,
		CommissionRequestCount: filter.CommissionRequestCount,
		CommissionAcceptCount:  filter.CommissionAcceptCount,
		CommissionSuccessCount: filter.CommissionSuccessCount,
		AvgRatings:             filter.AvgRatings,
		LastRequestTime:        filter.LastRequestTime,
	}
}

type SearchArtistsFilter struct {
	State                  *[]model.UserState `json:"state,omitempty"`
	RegTime                *model.TimeRange   `json:"reg_time,omitempty"`
	PaymentMethods         *[]string          `json:"payment_methods,omitempty"`
	YearOfDrawing          *model.TimeRange   `json:"year_of_drawing,omitempty"`
	CommissionRequestCount *model.IntRange    `json:"commission_request_count,omitempty"`
	CommissionAcceptCount  *model.IntRange    `json:"commission_accept_count,omitempty"`
	CommissionSuccessCount *model.IntRange    `json:"commission_success_count,omitempty"`
	AvgRatings             *model.FloatRange  `json:"avg_ratings,omitempty"`
	LastRequestTime        *model.TimeRange   `json:"last_request_time,omitempty"`
}

func getSorter(sorter model.ArtistSorter) SearchArtistsSorter {
	return SearchArtistsSorter{
		ArtistID:               sorter.ArtistID,
		UserName:               sorter.UserName,
		RegTime:                sorter.RegTime,
		YearOfDrawing:          sorter.YearOfDrawing,
		CommissionRequestCount: sorter.CommissionRequestCount,
		CommissionAcceptCount:  sorter.CommissionAcceptCount,
		CommissionSuccessCount: sorter.CommissionSuccessCount,
		AvgRatings:             sorter.AvgRatings,
		LastRequestTime:        sorter.LastRequestTime,
	}
}

type SearchArtistsSorter struct {
	ArtistID               *model2.SortOrder
	UserName               *model2.SortOrder
	RegTime                *model2.SortOrder
	YearOfDrawing          *model2.SortOrder
	CommissionRequestCount *model2.SortOrder
	CommissionAcceptCount  *model2.SortOrder
	CommissionSuccessCount *model2.SortOrder
	AvgRatings             *model2.SortOrder
	LastRequestTime        *model2.SortOrder
}
