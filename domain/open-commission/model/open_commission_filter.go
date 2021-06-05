package model

import "pixstall-search/domain/model"

type OpenCommissionFilter struct {
	State *[]OpenCommissionState

	PriceAmount   *model.FloatRange
	PriceCurrency *Currency

	DayNeed    *model.IntRange
	IsR18          *bool
	AllowBePrivate *bool
	AllowAnonymous *bool

	PageFilter model.PageFilter
}
