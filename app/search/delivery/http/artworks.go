package http

import (
	"pixstall-search/domain/artwork/model"
	model3 "pixstall-search/domain/model"
)

func getArtworkSorter(str string) *model.ArtworkSorter {
	if str == "" {
		return nil
	}
	sorter := model.ArtworkSorter{}
	symbol := str[:1]
	if symbol == "-" {
		switch str[1:len(str)] {
		case "day-used":
			v := model3.SortOrderDescending
			sorter.DayUsed = &v
			return &sorter
		case "rating":
			v := model3.SortOrderDescending
			sorter.Rating = &v
			return &sorter
		case "views":
			v := model3.SortOrderDescending
			sorter.Views = &v
			return &sorter
		case "favor-count":
			v := model3.SortOrderDescending
			sorter.FavorCount = &v
			return &sorter
		case "completed-time":
			v := model3.SortOrderDescending
			sorter.CompletedTime = &v
			return &sorter
		}
	} else {
		switch str {
		case "day-used":
			v := model3.SortOrderAscending
			sorter.DayUsed = &v
			return &sorter
		case "rating":
			v := model3.SortOrderAscending
			sorter.Rating = &v
			return &sorter
		case "views":
			v := model3.SortOrderAscending
			sorter.Views = &v
			return &sorter
		case "favor-count":
			v := model3.SortOrderAscending
			sorter.FavorCount = &v
			return &sorter
		case "completed-time":
			v := model3.SortOrderAscending
			sorter.CompletedTime = &v
			return &sorter
		}
	}
	return nil
}