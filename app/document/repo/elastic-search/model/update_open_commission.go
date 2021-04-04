package model

import (
	"pixstall-search/domain/open-commission/model"
	"time"
)

type UpdateOpenCommissionRequest struct {
	ID               string                     `json:"id"`
	Title            *string                    `json:"title,omitempty"`
	Desc             *string                    `json:"desc,omitempty"`
	ConvPrice        *float64                   `json:"conv_price,omitempty"`
	PriceAmount      *float64                   `json:"price_amount,omitempty"`
	PriceCurrency    *model.Currency            `json:"price_currency,omitempty"`
	DayNeedFrom      *int                       `json:"day_need_from,omitempty"`
	DayNeedTo        *int                       `json:"day_need_to,omitempty"`
	SampleImagePaths *[]string                  `json:"sample_image_paths,omitempty"`
	IsR18            *bool                      `json:"is_r18,omitempty"`
	AllowBePrivate   *bool                      `json:"allow_be_private,omitempty"`
	AllowAnonymous   *bool                      `json:"allow_anonymous,omitempty"`
	State            *model.OpenCommissionState `json:"state,omitempty"`
	LastUpdatedTime  time.Time                  `json:"last_updated_time"`
}

func NewUpdateOpenCommissionRequest(updater model.OpenCommissionUpdater) UpdateOpenCommissionRequest {

	var convPrice *float64
	var priceAmount *float64
	var priceCurrency *model.Currency
	if updater.Price != nil {
		cp := updater.Price.GetConvPrice()
		convPrice = &cp
		priceAmount = &updater.Price.Amount
		priceCurrency = &updater.Price.Currency
	}
	var dayNeedFrom *int
	var dayNeedTo *int
	if updater.DayNeed != nil {
		dayNeedFrom = &updater.DayNeed.From
		dayNeedTo = &updater.DayNeed.To
	}
	return UpdateOpenCommissionRequest{
		ID:               updater.ID,
		Title:            updater.Title,
		Desc:             updater.Desc,
		ConvPrice:        convPrice,
		PriceAmount:      priceAmount,
		PriceCurrency:    priceCurrency,
		DayNeedFrom:      dayNeedFrom,
		DayNeedTo:        dayNeedTo,
		SampleImagePaths: updater.SampleImagePaths,
		IsR18:            updater.IsR18,
		AllowBePrivate:   updater.AllowBePrivate,
		AllowAnonymous:   updater.AllowAnonymous,
		State:            updater.State,
		LastUpdatedTime:  updater.LastUpdatedTime,
	}
}
