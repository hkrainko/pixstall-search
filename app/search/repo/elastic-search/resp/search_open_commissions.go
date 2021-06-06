package resp

import (
	model2 "pixstall-search/domain/model"
	"pixstall-search/domain/open-commission/model"
	"strconv"
	"time"
)

type SearchOpenCommissionsResponse struct {
	Meta struct {
		Page Page `json:"page"`
	} `json:"meta"`
	Results []SearchOpenCommissionResponseResult `json:"results"`
}

type SearchOpenCommissionResponseResult struct {
	ID struct {
		Raw string `json:"raw"`
	} `json:"id"`
	ArtistID struct {
		Raw string `json:"raw"`
	} `json:"artist_id"`
	Title struct {
		Raw string `json:"raw"`
	} `json:"title"`
	Desc struct {
		Raw string `json:"raw"`
	} `json:"desc"`

	// Price
	PriceAmount struct {
		Raw float64 `json:"raw"`
	} `json:"price_amount"`
	PriceCurrency struct {
		Raw model.Currency `json:"raw"`
	} `json:"price_currency"`

	// DayNeed
	DayNeedForm struct {
		Raw float64 `json:"raw"`
	} `json:"day_need_from"`
	DayNeedTo struct {
		Raw float64 `json:"raw"`
	} `json:"day_need_to"`

	SampleImagePaths struct {
		Raw []string `json:"raw"`
	} `json:"sample_image_paths"`
	IsR18 struct {
		Raw string `json:"raw"`
	} `json:"is_r18"`
	AllowBePrivate struct {
		Raw string `json:"raw"`
	} `json:"allow_be_private"`
	AllowAnonymous struct {
		Raw string `json:"raw"`
	} `json:"allow_anonymous"`
	State struct {
		Raw model.OpenCommissionState `json:"raw"`
	} `json:"state"`
	CreateTime struct {
		Raw time.Time `json:"raw"`
	} `json:"create_time"`
	LastUpdatedTime struct {
		Raw time.Time `json:"raw"`
	} `json:"last_updated_time"`
}

func (r SearchOpenCommissionsResponse) ToDomainResult() model.GetOpenCommissionsResult {
	openCommissions := make([]model.OpenCommission, 0)
	for _, v := range r.Results {
		isR18, err := strconv.ParseBool(v.IsR18.Raw)
		if err != nil {
			continue
		}
		allowBePrivate, err := strconv.ParseBool(v.AllowBePrivate.Raw)
		if err != nil {
			continue
		}
		allowAnonymous, err := strconv.ParseBool(v.AllowAnonymous.Raw)
		if err != nil {
			continue
		}
		openCommissions = append(openCommissions, model.OpenCommission{
			ID:       v.ID.Raw,
			ArtistID: v.ArtistID.Raw,
			Title:    v.Title.Raw,
			Desc:     v.Desc.Raw,
			Price: model.Price{
				Amount:   v.PriceAmount.Raw,
				Currency: v.PriceCurrency.Raw,
			},
			DayNeed: model.DayNeed{
				From: int(v.DayNeedForm.Raw),
				To:   int(v.DayNeedTo.Raw),
			},
			SampleImagePaths: v.SampleImagePaths.Raw,
			IsR18:            isR18,
			AllowBePrivate:   allowBePrivate,
			AllowAnonymous:   allowAnonymous,
			State:            v.State.Raw,
			CreateTime:       v.CreateTime.Raw,
			LastUpdatedTime:  v.LastUpdatedTime.Raw,
		})
	}
	return model.GetOpenCommissionsResult{
		Page: model2.Page{
			Current:     r.Meta.Page.Current,
			TotalPage:   r.Meta.Page.TotalPage,
			TotalResult: r.Meta.Page.TotalResult,
			Size:        r.Meta.Page.Size,
		},
		OpenCommissions: openCommissions,
	}
}
