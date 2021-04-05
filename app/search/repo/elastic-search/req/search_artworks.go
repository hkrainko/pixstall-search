package req

import (
	"pixstall-search/domain/artwork/model"
	model2 "pixstall-search/domain/model"
)

type SearchArtworksRequest struct {
	Query  string                `json:"query"`
	Page   model2.PageFilter     `json:"page"`
	Filter *ESFilters            `json:"filters,omitempty"`
	Sort   *SearchArtworksSorter `json:"sort,omitempty"`
}

func NewSearchArtworksRequest(query string, filter model.ArtworkFilter, sorter model.ArtworkSorter) SearchArtworksRequest {
	var filters *ESFilters
	if fs := getSearchArtworksFilter(filter); len(fs) > 0 {
		filters = &ESFilters{
			All: fs,
		}
	}
	s := getSearchArtworksSorter(sorter)
	return SearchArtworksRequest{
		Query:  query,
		Page:   filter.PageFilter,
		Filter: filters,
		Sort:   &s,
	}
}

func getSearchArtworksFilter(filter model.ArtworkFilter) []map[string]interface{} {
	var filters []map[string]interface{}
	if filter.State != nil {
		filters = append(filters, map[string]interface{}{"state": filter.State})
	}
	if filter.DayUsed != nil {
		filters = append(filters, map[string]interface{}{"day_used": filter.DayUsed})
	}
	if filter.IsR18 != nil {
		filters = append(filters, map[string]interface{}{"is_r18": filter.IsR18})
	}
	if filter.Anonymous != nil {
		filters = append(filters, map[string]interface{}{"anonymous": filter.Anonymous})
	}
	if filter.Rating != nil {
		filters = append(filters, map[string]interface{}{"rating": filter.Rating})
	}
	if filter.Views != nil {
		filters = append(filters, map[string]interface{}{"views": filter.Views})
	}
	if filter.FavorCount != nil {
		filters = append(filters, map[string]interface{}{"favor_count": filter.FavorCount})
	}
	if filter.CreateTime != nil {
		filters = append(filters, map[string]interface{}{"create_time": filter.CreateTime})
	}
	if filter.StartTime != nil {
		filters = append(filters, map[string]interface{}{"start_time": filter.StartTime})
	}
	if filter.CompletedTime != nil {
		filters = append(filters, map[string]interface{}{"completed_time": filter.CompletedTime})
	}
	if filter.LastUpdatedTime != nil {
		filters = append(filters, map[string]interface{}{"last_update_time": filter.LastUpdatedTime})
	}
	return filters
}

func getSearchArtworksSorter(sorter model.ArtworkSorter) SearchArtworksSorter {
	if sorter.ArtistID == nil &&
		sorter.UserName == nil &&
		sorter.DayUsed == nil &&
		sorter.Rating == nil &&
		sorter.Views == nil &&
		sorter.FavorCount == nil &&
		sorter.CreateTime == nil &&
		sorter.StartTime == nil &&
		sorter.CompletedTime == nil &&
		sorter.LastUpdatedTime == nil {
		desc := model2.SortOrderDescending
		return SearchArtworksSorter{
			Score: &desc,
		}
	}

	return SearchArtworksSorter{
		ArtistID:        sorter.ArtistID,
		UserName:        sorter.UserName,
		DayUsed:         sorter.DayUsed,
		Rating:          sorter.Rating,
		Views:           sorter.Views,
		FavorCount:      sorter.FavorCount,
		CreateTime:      sorter.CreateTime,
		StartTime:       sorter.StartTime,
		CompletedTime:   sorter.CompletedTime,
		LastUpdatedTime: sorter.LastUpdatedTime,
	}
}

type SearchArtworksSorter struct {
	Score           *model2.SortOrder `json:"_score,omitempty"`
	ArtistID        *model2.SortOrder `json:"artist_id,omitempty"`
	UserName        *model2.SortOrder `json:"user_name,omitempty"`
	DayUsed         *model2.SortOrder `json:"day_used,omitempty"`
	Rating          *model2.SortOrder `json:"rating,omitempty"`
	Views           *model2.SortOrder `json:"views,omitempty"`
	FavorCount      *model2.SortOrder `json:"favor_count,omitempty"`
	CreateTime      *model2.SortOrder `json:"create_time,omitempty"`
	StartTime       *model2.SortOrder `json:"start_time,omitempty"`
	CompletedTime   *model2.SortOrder `json:"completed_time,omitempty"`
	LastUpdatedTime *model2.SortOrder `json:"last_updated_time,omitempty"`
}
