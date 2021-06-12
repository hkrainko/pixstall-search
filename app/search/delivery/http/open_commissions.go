package http

import (
	model3 "pixstall-search/domain/model"
	model2 "pixstall-search/domain/open-commission/model"
)

func getOpenCommissionSorter(str string) *model2.OpenCommissionSorter {
	if str == "" {
		return nil
	}
	sorter := model2.OpenCommissionSorter{}
	symbol := str[:1]
	if symbol == "-" {
		switch str[1:len(str)] {
		case "artist-id":
			v := model3.SortOrderDescending
			sorter.ArtistID = &v
			return &sorter
		case "price-from":
			v := model3.SortOrderDescending
			sorter.Price = &v
			return &sorter
		case "day-need-from":
			v := model3.SortOrderDescending
			sorter.DayNeedFrom = &v
			return &sorter
		case "day-need-to":
			v := model3.SortOrderDescending
			sorter.DayNeedTo = &v
			return &sorter
		case "create-time":
			v := model3.SortOrderDescending
			sorter.CreateTime = &v
			return &sorter
		case "last-updated-time":
			v := model3.SortOrderDescending
			sorter.LastUpdatedTime = &v
			return &sorter
		}
	} else {
		switch str {
		case "artist-id":
			v := model3.SortOrderAscending
			sorter.ArtistID = &v
			return &sorter
		case "price-from":
			v := model3.SortOrderAscending
			sorter.Price = &v
			return &sorter
		case "day-need-from":
			v := model3.SortOrderAscending
			sorter.DayNeedFrom = &v
			return &sorter
		case "day-need-to":
			v := model3.SortOrderAscending
			sorter.DayNeedTo = &v
			return &sorter
		case "create-time":
			v := model3.SortOrderAscending
			sorter.CreateTime = &v
			return &sorter
		case "last-updated-time":
			v := model3.SortOrderAscending
			sorter.LastUpdatedTime = &v
			return &sorter
		}
	}
	return nil
}