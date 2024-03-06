package repository

import (
	"github.com/arvan-challenge/src/entity"
	"gorm.io/gorm"
)

type chargeCodeRepository struct {
	CommonBehaviourRepository[entity.ChargeCode]
}

func NewChargeCodeRepository(db *gorm.DB) ChargeCodeRepository {
	return &chargeCodeRepository{
		CommonBehaviourRepository: NewCommonBehaviour[entity.ChargeCode](db),
	}
}
