package repository

import (
	"github.com/arvan-challenge/src/entity"
	"gorm.io/gorm"
)

type transactionRepository struct {
	CommonBehaviourRepository[entity.Transaction]
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		CommonBehaviourRepository: NewCommonBehaviour[entity.Transaction](db),
	}
}
