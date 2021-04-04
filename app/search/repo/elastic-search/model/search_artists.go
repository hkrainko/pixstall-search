package model

import (
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/model"
)

type SearchArtistsRequest struct {
	Query   string               `json:"query"`
	Page    model2.PageFilter    `json:"page"`
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
		Page:    filter.PageFilter,
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
	if sorter.ArtistID == nil &&
		sorter.UserName == nil &&
		sorter.RegTime == nil &&
		sorter.YearOfDrawing == nil &&
		sorter.CommissionRequestCount == nil &&
		sorter.CommissionAcceptCount == nil &&
		sorter.CommissionSuccessCount == nil &&
		sorter.AvgRatings == nil &&
		sorter.LastRequestTime == nil {
		desc := model2.SortOrderDescending
		return SearchArtistsSorter{
			Score: &desc,
		}
	}
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
	Score                  *model2.SortOrder `json:"_score,omitempty"`
	ArtistID               *model2.SortOrder `json:"artist_id,omitempty"`
	UserName               *model2.SortOrder `json:"user_name,omitempty"`
	RegTime                *model2.SortOrder `json:"reg_time,omitempty"`
	YearOfDrawing          *model2.SortOrder `json:"year_of_drawing,omitempty"`
	CommissionRequestCount *model2.SortOrder `json:"commission_request_count,omitempty"`
	CommissionAcceptCount  *model2.SortOrder `json:"commission_accept_count,omitempty"`
	CommissionSuccessCount *model2.SortOrder `json:"commission_success_count,omitempty"`
	AvgRatings             *model2.SortOrder `json:"avg_ratings,omitempty"`
	LastRequestTime        *model2.SortOrder `json:"last_request_time,omitempty"`
}
