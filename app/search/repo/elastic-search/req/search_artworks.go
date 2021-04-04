package req

import (
	"pixstall-search/domain/artwork/model"
	model2 "pixstall-search/domain/model"
	"time"
)

type SearchArtworksRequest struct {
	Query  string                `json:"query"`
	Page   model2.PageFilter     `json:"page"`
	Filter *SearchArtworksFilter `json:"filters,omitempty"`
	Sort   *SearchArtworksSorter `json:"sort,omitempty"`
}

func NewSearchArtworksRequest(query string, filter model.ArtworkFilter, sorter model.ArtworkSorter) SearchArtworksRequest {
	f := getSearchArtworksFilter(filter)
	s := getSearchArtworksSorter(sorter)
	return SearchArtworksRequest{
		Query:  query,
		Page:   filter.PageFilter,
		Filter: &f,
		Sort:   &s,
	}
}

func getSearchArtworksFilter(filter model.ArtworkFilter) SearchArtworksFilter {
	return SearchArtworksFilter{
		State:          filter.State,
		DayUsed:        filter.DayUsed,
		IsR18:          filter.IsR18,
		Anonymous:      filter.Anonymous,
		Rating:         filter.Rating,
		Views:          filter.Views,
		FavorCount:     filter.FavorCount,
		CreateTime:     filter.CreateTime,
		StartTime:      filter.StartTime,
		CompletedTime:  filter.CompletedTime,
		LastUpdateTime: filter.LastUpdateTime,
	}
}

type SearchArtworksFilter struct {
	State          *[]model.ArtworkState `json:"state,omitempty"`
	DayUsed        *time.Duration        `json:"day_used,omitempty"`
	IsR18          *bool                 `json:"is_r18,omitempty"`
	Anonymous      *bool                 `json:"anonymous,omitempty"`
	Rating         *model2.IntRange      `json:"rating,omitempty"`
	Views          *model2.IntRange      `json:"views,omitempty"`
	FavorCount     *model2.IntRange      `json:"favor_count,omitempty"`
	CreateTime     *model2.TimeRange     `json:"create_time,omitempty"`
	StartTime      *model2.TimeRange     `json:"start_time,omitempty"`
	CompletedTime  *model2.TimeRange     `json:"completed_time,omitempty"`
	LastUpdateTime *model2.TimeRange     `json:"last_update_time,omitempty"`
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
		sorter.LastUpdateTime == nil {
		desc := model2.SortOrderDescending
		return SearchArtworksSorter{
			Score: &desc,
		}
	}

	return SearchArtworksSorter{
		ArtistID:       sorter.ArtistID,
		UserName:       sorter.UserName,
		DayUsed:        sorter.DayUsed,
		Rating:         sorter.Rating,
		Views:          sorter.Views,
		FavorCount:     sorter.FavorCount,
		CreateTime:     sorter.CreateTime,
		StartTime:      sorter.StartTime,
		CompletedTime:  sorter.CompletedTime,
		LastUpdateTime: sorter.LastUpdateTime,
	}
}

type SearchArtworksSorter struct {
	Score          *model2.SortOrder `json:"_score,omitempty"`
	ArtistID       *model2.SortOrder `json:"artist_id,omitempty"`
	UserName       *model2.SortOrder `json:"user_name,omitempty"`
	DayUsed        *model2.SortOrder `json:"day_used,omitempty"`
	Rating         *model2.SortOrder `json:"rating,omitempty"`
	Views          *model2.SortOrder `json:"views,omitempty"`
	FavorCount     *model2.SortOrder `json:"favor_count,omitempty"`
	CreateTime     *model2.SortOrder `json:"create_time,omitempty"`
	StartTime      *model2.SortOrder `json:"start_time,omitempty"`
	CompletedTime  *model2.SortOrder `json:"completed_time,omitempty"`
	LastUpdateTime *model2.SortOrder `json:"last_update_time,omitempty"`
}
