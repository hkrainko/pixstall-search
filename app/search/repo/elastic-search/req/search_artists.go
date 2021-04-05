package req

import (
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/model"
)

type SearchArtistsRequest struct {
	Query   string               `json:"query"`
	Page    model2.PageFilter    `json:"page"`
	Filters *ESFilters           `json:"filters,omitempty"`
	Sort    *SearchArtistsSorter `json:"sort,omitempty"`
}

func NewSearchArtistRequest(query string, filter model.ArtistFilter, sorter model.ArtistSorter) SearchArtistsRequest {
	var filters *ESFilters
	if fs := getSearchArtistsFilters(filter); len(fs) > 0 {
		filters = &ESFilters{
			All: fs,
		}
	}
	s := getSearchArtistsSorter(sorter)
	return SearchArtistsRequest{
		Query:   query,
		Page:    filter.PageFilter,
		Filters: filters,
		Sort:    &s,
	}
}

func getSearchArtistsFilters(filter model.ArtistFilter) []map[string]interface{} {
	var filters []map[string]interface{}
	if filter.State != nil {
		filters = append(filters, map[string]interface{}{"state": filter.State})
	}
	if filter.RegTime != nil {
		filters = append(filters, map[string]interface{}{"reg_time": filter.RegTime})
	}
	if filter.PaymentMethods != nil {
		filters = append(filters, map[string]interface{}{"payment_methods": filter.PaymentMethods})
	}
	if filter.YearOfDrawing != nil {
		filters = append(filters, map[string]interface{}{"year_of_drawing": filter.YearOfDrawing})
	}
	if filter.CommissionRequestCount != nil {
		filters = append(filters, map[string]interface{}{"commission_request_count": filter.CommissionRequestCount})
	}
	if filter.CommissionAcceptCount != nil {
		filters = append(filters, map[string]interface{}{"commission_accept_count": filter.CommissionAcceptCount})
	}
	if filter.CommissionSuccessCount != nil {
		filters = append(filters, map[string]interface{}{"commission_success_count": filter.CommissionSuccessCount})
	}
	if filter.AvgRatings != nil {
		filters = append(filters, map[string]interface{}{"avg_ratings": filter.AvgRatings})
	}
	if filter.LastRequestTime != nil {
		filters = append(filters, map[string]interface{}{"last_request_time": filter.LastRequestTime})
	}
	return filters
}

func getSearchArtistsSorter(sorter model.ArtistSorter) SearchArtistsSorter {
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
