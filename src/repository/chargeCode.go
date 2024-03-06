package repository

import (
	"github.com/arvan-challenge/src/entity"
	repository "github.com/arvan-challenge/src/repository/interface"
	"gorm.io/gorm"
)

type chargeCodeRepository struct {
	repository.CommonBehaviourRepository[entity.ChargeCode]
}

func NewChargeCodeRepository(db *gorm.DB) repository.ChargeCodeRepository {
	return &chargeCodeRepository{
		CommonBehaviourRepository: repository.NewCommonBehaviour[entity.ChargeCode](db),
	}
}
