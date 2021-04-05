package model

import (
	"pixstall-search/domain/model"
)

type ArtworkSorter struct {
	ArtistID        *model.SortOrder
	UserName        *model.SortOrder
	DayUsed         *model.SortOrder
	Rating          *model.SortOrder
	Views           *model.SortOrder
	FavorCount      *model.SortOrder
	CreateTime      *model.SortOrder
	StartTime       *model.SortOrder
	CompletedTime   *model.SortOrder
	LastUpdatedTime *model.SortOrder
}
