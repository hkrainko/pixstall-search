package model

import (
	"pixstall-search/domain/model"
	"time"
)

type ArtworkFilter struct {
	State *[]ArtworkState

	DayUsed   *time.Duration
	IsR18     *bool
	Anonymous *bool

	Rating *model.IntRange

	Views      *model.IntRange
	FavorCount *model.IntRange

	CreateTime     *model.TimeRange
	StartTime      *model.TimeRange
	CompletedTime  *model.TimeRange
	LastUpdateTime *model.TimeRange

	PageFilter model.PageFilter
}
