package open_commission

import (
	"context"
	"pixstall-search/domain/open-commission/model"
)

type UseCase interface {
	CreateOpenCommission(ctx context.Context, creator model.OpenCommissionCreator) (*string, error)
	UpdateOpenCommission(ctx context.Context, updater model.OpenCommissionUpdater) (*string, error)
}
