package usecase

import (
	"context"
	"pixstall-search/domain/document"
	open_commission "pixstall-search/domain/open-commission"
	"pixstall-search/domain/open-commission/model"
)

type openCommissionUseCase struct {
	docRepo document.Repo
}

func NewOpenCommissionUseCase(docRepo document.Repo) open_commission.UseCase {
	return &openCommissionUseCase{
		docRepo: docRepo,
	}
}

func (o openCommissionUseCase) CreateOpenCommission(ctx context.Context, creator model.OpenCommissionCreator) (*string, error) {
	return o.docRepo.AddOpenCommission(ctx, creator)
}

func (o openCommissionUseCase) UpdateOpenCommission(ctx context.Context, updater model.OpenCommissionUpdater) (*string, error) {
	return o.docRepo.UpdateOpenCommission(ctx, updater)
}
