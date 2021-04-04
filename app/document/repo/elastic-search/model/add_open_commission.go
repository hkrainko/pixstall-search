package model

import (
	"pixstall-search/domain/open-commission/model"
	"time"
)

type AddOpenCommissionRequest struct {
	ID       string `json:"id"`
	ArtistID string `json:"artist_id"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`

	// Price
	PriceAmount   float64        `json:"price_amount"`
	PriceCurrency model.Currency `json:"price_currency"`

	// DayNeed
	DayNeedForm int `json:"day_need_from"`
	DayNeedTo   int `json:"day_need_to"`

	SampleImagePaths []string                  `json:"sample_image_paths"`
	IsR18            bool                      `json:"is_r18"`
	AllowBePrivate   bool                      `json:"allow_be_private"`
	AllowAnonymous   bool                      `json:"allow_anonymous"`
	State            model.OpenCommissionState `json:"state"`
	CreateTime       time.Time                 `json:"create_time"`
	LastUpdatedTime  time.Time                 `json:"last_updated_time"`
}

func NewAddOpenCommissionRequest(creator model.OpenCommissionCreator) AddOpenCommissionRequest {
	return AddOpenCommissionRequest{
		ID:               creator.ID,
		ArtistID:         creator.ArtistID,
		Title:            creator.Title,
		Desc:             creator.Desc,
		PriceAmount:      creator.Price.Amount,
		PriceCurrency:    creator.Price.Currency,
		DayNeedForm:      creator.DayNeed.From,
		DayNeedTo:        creator.DayNeed.To,
		SampleImagePaths: creator.SampleImagePaths,
		IsR18:            creator.IsR18,
		AllowBePrivate:   creator.AllowBePrivate,
		AllowAnonymous:   creator.AllowAnonymous,
		State:            creator.State,
		CreateTime:       creator.CreateTime,
		LastUpdatedTime:  creator.LastUpdatedTime,
	}
}
