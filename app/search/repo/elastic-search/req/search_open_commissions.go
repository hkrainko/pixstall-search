package req

import (
	model2 "pixstall-search/domain/model"
	"pixstall-search/domain/open-commission/model"
)

type SearchOpenCommissionsRequest struct {
	Query  string                       `json:"query"`
	Page   model2.PageFilter            `json:"page"`
	Filter *SearchOpenCommissionsFilter `json:"filters,omitempty"`
	Sort   *SearchOpenCommissionsSorter `json:"sort,omitempty"`
}

func NewSearchOpenCommissionsRequest(query string, filter model.OpenCommissionFilter, sorter model.OpenCommissionSorter) SearchOpenCommissionsRequest {
	f := getSearchOpenCommissionFilter(filter)
	s := getSearchOpenCommissionsSorter(sorter)
	return SearchOpenCommissionsRequest{
		Query:  query,
		Page:   filter.PageFilter,
		Filter: &f,
		Sort:   &s,
	}
}

func getSearchOpenCommissionFilter(filter model.OpenCommissionFilter) SearchOpenCommissionsFilter {
	var convPrice *model2.FloatRange
	if filter.PriceAmount != nil && filter.PriceCurrency != nil {
		from := filter.PriceCurrency.GetConvPrice(*filter.PriceAmount.From)
		to := filter.PriceCurrency.GetConvPrice(*filter.PriceAmount.To)
		convPrice = &model2.FloatRange{
			From: &from,
			To:   &to,
		}
	}
	return SearchOpenCommissionsFilter{
		State:          filter.State,
		ConvPrice:      convPrice,
		DayNeedFrom:    filter.DayNeedFrom,
		DayNeedTo:      filter.DayNeedTo,
		IsR18:          filter.IsR18,
		AllowBePrivate: filter.AllowBePrivate,
		AllowAnonymous: filter.AllowAnonymous,
	}
}

type SearchOpenCommissionsFilter struct {
	State          *[]model.OpenCommissionState `json:"state,omitempty"`
	ConvPrice      *model2.FloatRange           `json:"conv_price,omitempty"`
	DayNeedFrom    *model2.IntRange             `json:"day_need_from,omitempty"`
	DayNeedTo      *model2.IntRange             `json:"day_need_to,omitempty"`
	IsR18          *bool                        `json:"is_r18,omitempty"`
	AllowBePrivate *bool                        `json:"allow_be_private,omitempty"`
	AllowAnonymous *bool                        `json:"allow_anonymous,omitempty"`
}

func getSearchOpenCommissionsSorter(sorter model.OpenCommissionSorter) SearchOpenCommissionsSorter {
	if sorter.ArtistID == nil &&
		sorter.Price == nil &&
		sorter.DayNeedFrom == nil &&
		sorter.DayNeedTo == nil &&
		sorter.CreateTime == nil &&
		sorter.LastUpdatedTime == nil {
		desc := model2.SortOrderDescending
		return SearchOpenCommissionsSorter{
			Score: &desc,
		}
	}
	return SearchOpenCommissionsSorter{
		ArtistID:        sorter.ArtistID,
		Price:           sorter.Price,
		DayNeedFrom:     sorter.DayNeedFrom,
		DayNeedTo:       sorter.DayNeedTo,
		CreateTime:      sorter.CreateTime,
		LastUpdatedTime: sorter.LastUpdatedTime,
	}
}

type SearchOpenCommissionsSorter struct {
	Score           *model2.SortOrder `json:"_score,omitempty"`
	ArtistID        *model2.SortOrder `json:"artist_id,omitempty"`
	Price           *model2.SortOrder `json:"price,omitempty"`
	DayNeedFrom     *model2.SortOrder `json:"day_need_from,omitempty"`
	DayNeedTo       *model2.SortOrder `json:"day_need_to,omitempty"`
	CreateTime      *model2.SortOrder `json:"create_time,omitempty"`
	LastUpdatedTime *model2.SortOrder `json:"last_updated_time,omitempty"`
}
