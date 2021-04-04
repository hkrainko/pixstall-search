package model

import "pixstall-search/domain/model"

type GetOpenCommissionsResult struct {
	Page            model.Page       `json:"page"`
	OpenCommissions []OpenCommission `json:"open_commissions"`
}
