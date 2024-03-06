package repository

import (
	"github.com/arvan-challenge/src/entity"
	repository "github.com/arvan-challenge/src/repository/interface"
	"gorm.io/gorm"
)

type disCountCodeRepository struct {
	repository.CommonBehaviourRepository[entity.DiscountCode]
}

func NewDisCountCodeRepository(db *gorm.DB) repository.DiscountCodeRepository {
	return &disCountCodeRepository{
		CommonBehaviourRepository: repository.NewCommonBehaviour[entity.DiscountCode](db),
	}
}
