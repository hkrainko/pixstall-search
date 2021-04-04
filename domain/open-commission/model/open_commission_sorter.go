package model

import "pixstall-search/domain/model"

type OpenCommissionSorter struct {
	ArtistID        *model.SortOrder
	Price           *model.SortOrder
	DayNeedFrom     *model.SortOrder
	DayNeedTo       *model.SortOrder
	CreateTime      *model.SortOrder
	LastUpdatedTime *model.SortOrder
}
