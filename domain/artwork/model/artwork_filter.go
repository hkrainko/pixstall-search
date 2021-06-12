package model

import (
	"pixstall-search/domain/model"
)

type ArtworkFilter struct {
	State *[]ArtworkState

	DayUsed   *model.IntRange
	IsR18     *bool
	Anonymous *bool

	Rating *model.IntRange

	Views      *model.IntRange
	FavorCount *model.IntRange

	CreateTime      *model.TimeRange
	StartTime       *model.TimeRange
	CompletedTime   *model.TimeRange
	LastUpdatedTime *model.TimeRange

	PageFilter model.PageFilter
}
