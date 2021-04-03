package usecase

import (
	"context"
	open_commission "pixstall-search/domain/open-commission"
	"pixstall-search/domain/open-commission/model"
)

type openCommissionUseCase struct {

}

func NewOpenCommissionUseCase() open_commission.UseCase {
	return &openCommissionUseCase{}
}

func (o openCommissionUseCase) CreateOpenCommission(ctx context.Context, creator model.OpenCommissionCreator) (*string, error) {
	panic("implement me")
}

func (o openCommissionUseCase) UpdateOpenCommission(ctx context.Context, updater model.OpenCommissionUpdater) (*string, error) {
	panic("implement me")
}