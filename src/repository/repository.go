package repository

import (
	"context"
	"github.com/arvan-challenge/src/entity"
	"github.com/arvan-challenge/utils"
)

type CommonBehaviourRepository[T utils.DBModel] interface {
	ByID(ctx context.Context, id uint) (T, error)
	ByField(ctx context.Context, field string, id uint) (T, error)
	Save(ctx context.Context, model *T) error
	// add more common behaviour
}

type UserRepository interface {
	CommonBehaviourRepository[entity.User]
}

type ChargeCodeRepository interface {
	CommonBehaviourRepository[entity.ChargeCode]
}

type DiscountCodeRepository interface {
	CommonBehaviourRepository[entity.DiscountCode]
}

type TransactionRepository interface {
	CommonBehaviourRepository[entity.Transaction]
}
