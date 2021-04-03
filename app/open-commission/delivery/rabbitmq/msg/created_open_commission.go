package msg

import (
	"pixstall-search/domain/open-commission/model"
	"time"
)

type CreatedOpenCommission struct {
	ID                             string                   `json:"id"`
	ArtistID                       string                   `json:"artistId"`
	Title                          string                   `json:"title"`
	Desc                           string                   `json:"desc"`
	DepositRule                    *string                  `json:"depositRule"`
	Price                          model.Price              `json:"price"`
	DayNeed                        model.DayNeed            `json:"dayNeed"`
	TimesAllowedDraftToChange      *int                     `json:"timesAllowedDraftToChange"`
	TimesAllowedCompletionToChange *int                     `json:"timesAllowedCompletionToChange"`
	SampleImagePaths               []string                 `json:"sampleImagePaths"`
	IsR18                          bool                     `json:"isR18"`
	AllowBePrivate                 bool                     `json:"allowBePrivate"`
	AllowAnonymous                 bool                     `json:"allowAnonymous"`
	State                          model.OpenCommissionSate `json:"state"`
	CreateTime                     time.Time                `json:"createTime"`
	LastUpdatedTime                time.Time                `json:"lastUpdatedTime"`
}

func (c CreatedOpenCommission) ToDomainOpenCommissionCreator() model.OpenCommissionCreator {
	return model.OpenCommissionCreator{
		ID:               c.ID,
		ArtistID:         c.ArtistID,
		Title:            c.Title,
		Desc:             c.Desc,
		Price:            c.Price,
		DayNeed:          c.DayNeed,
		IsR18:            c.IsR18,
		AllowBePrivate:   c.AllowBePrivate,
		AllowAnonymous:   c.AllowAnonymous,
		SampleImagePaths: c.SampleImagePaths,
	}
}
