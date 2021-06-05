package search_open_commissions

import (
	model2 "pixstall-search/domain/model"
	"pixstall-search/domain/open-commission/model"
)

type Response struct {
	Page            model2.Page            `json:",inline"`
	OpenCommissions []model.OpenCommission `json:"openCommissions"`
}

func NewResponse(result model.GetOpenCommissionsResult) *Response {
	return &Response{
		Page:            result.Page,
		OpenCommissions: result.OpenCommissions,
	}
}
