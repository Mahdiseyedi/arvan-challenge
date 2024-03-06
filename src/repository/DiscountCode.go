package repository

import (
	"github.com/arvan-challenge/src/entity"
	"gorm.io/gorm"
)

type disCountCodeRepository struct {
	CommonBehaviourRepository[entity.DiscountCode]
}

func NewDisCountCodeRepository(db *gorm.DB) DiscountCodeRepository {
	return &disCountCodeRepository{
		CommonBehaviourRepository: NewCommonBehaviour[entity.DiscountCode](db),
	}
}
