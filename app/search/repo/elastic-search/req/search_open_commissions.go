package req

import (
	model2 "pixstall-search/domain/model"
	"pixstall-search/domain/open-commission/model"
)

type SearchOpenCommissionsRequest struct {
	Query   string                       `json:"query"`
	Page    model2.PageFilter            `json:"page"`
	Filters *ESFilters                   `json:"filters,omitempty"`
	Sort    *SearchOpenCommissionsSorter `json:"sort,omitempty"`
}

func NewSearchOpenCommissionsRequest(query string, filter model.OpenCommissionFilter, sorter model.OpenCommissionSorter) SearchOpenCommissionsRequest {
	var filters *ESFilters
	if fs := getSearchOpenCommissionFilter(filter); len(fs) > 0 {
		filters = &ESFilters{
			All: fs,
		}
	}
	s := getSearchOpenCommissionsSorter(sorter)
	return SearchOpenCommissionsRequest{
		Query:   query,
		Page:    filter.PageFilter,
		Filters: filters,
		Sort:    &s,
	}
}

func getSearchOpenCommissionFilter(filter model.OpenCommissionFilter) []map[string]interface{} {
	var filters []map[string]interface{}
	if filter.State != nil {
		filters = append(filters, map[string]interface{}{"state": filter.State})
	}
	if filter.PriceAmount != nil && filter.PriceCurrency != nil {
		from := filter.PriceCurrency.GetConvPrice(*filter.PriceAmount.From)
		to := filter.PriceCurrency.GetConvPrice(*filter.PriceAmount.To)
		convPrice := &model2.FloatRange{
			From: &from,
			To:   &to,
		}
		filters = append(filters, map[string]interface{}{"conv_price": convPrice})
	}
	if filter.DayNeed != nil {
		if filter.DayNeed != nil {
			filters = append(filters, map[string]interface{}{"day_need_to": filter.DayNeed})
		}
	}
	if filter.IsR18 != nil {
		if *filter.IsR18 {
			filters = append(filters, map[string]interface{}{"is_r18": "true"})
		} else {
			filters = append(filters, map[string]interface{}{"is_r18": "false"})
		}
	}
	if filter.AllowBePrivate != nil {
		if *filter.AllowBePrivate {
			filters = append(filters, map[string]interface{}{"allow_be_private": "true"})
		} else {
			filters = append(filters, map[string]interface{}{"allow_be_private": "false"})
		}
	}
	if filter.AllowAnonymous != nil {
		if *filter.AllowAnonymous {
			filters = append(filters, map[string]interface{}{"allow_anonymous": "true"})
		} else {
			filters = append(filters, map[string]interface{}{"allow_anonymous": "false"})
		}
	}
	return filters
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
