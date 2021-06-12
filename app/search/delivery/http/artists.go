package http

import (
	"pixstall-search/domain/artist/model"
	model3 "pixstall-search/domain/model"
)

func getArtistSorter(str string) *model.ArtistSorter {
	if str == "" {
		return nil
	}
	sorter := model.ArtistSorter{}
	symbol := str[:1]
	if symbol == "-" {
		switch str[1:len(str)] {
		case "user-name":
			v := model3.SortOrderDescending
			sorter.UserName = &v
			return &sorter
		case "reg-time":
			v := model3.SortOrderDescending
			sorter.RegTime = &v
			return &sorter
		case "commission-request-count":
			v := model3.SortOrderDescending
			sorter.CommissionRequestCount = &v
			return &sorter
		case "commission-accept-count":
			v := model3.SortOrderDescending
			sorter.CommissionAcceptCount = &v
			return &sorter
		case "commission-success-count":
			v := model3.SortOrderDescending
			sorter.CommissionSuccessCount = &v
			return &sorter
		case "avg-ratings":
			v := model3.SortOrderDescending
			sorter.AvgRatings = &v
			return &sorter
		case "last-request-time":
			v := model3.SortOrderDescending
			sorter.LastRequestTime = &v
			return &sorter
		}
	} else {
		switch str {
		case "user-name":
			v := model3.SortOrderAscending
			sorter.UserName = &v
			return &sorter
		case "reg-time":
			v := model3.SortOrderAscending
			sorter.RegTime = &v
			return &sorter
		case "commission-request-count":
			v := model3.SortOrderAscending
			sorter.CommissionRequestCount = &v
			return &sorter
		case "commission-accept-count":
			v := model3.SortOrderAscending
			sorter.CommissionAcceptCount = &v
			return &sorter
		case "commission-success-count":
			v := model3.SortOrderAscending
			sorter.CommissionSuccessCount = &v
			return &sorter
		case "avg-ratings":
			v := model3.SortOrderAscending
			sorter.AvgRatings = &v
			return &sorter
		case "last-request-time":
			v := model3.SortOrderAscending
			sorter.LastRequestTime = &v
			return &sorter
		}
	}
	return nil
}