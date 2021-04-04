package model

import "pixstall-search/domain/model"

type OpenCommissionFilter struct {
	State *[]OpenCommissionState

	PriceAmount   *model.FloatRange
	PriceCurrency *Currency

	DayNeedFrom    *model.TimeRange
	DayNeedTo      *model.TimeRange
	IsR18          *bool
	AllowBePrivate *bool
	AllowAnonymous *bool

	PageFilter model.PageFilter
}
