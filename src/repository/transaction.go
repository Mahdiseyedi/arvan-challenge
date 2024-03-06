package repository

import (
	"github.com/arvan-challenge/src/entity"
	repository "github.com/arvan-challenge/src/repository/interface"
	"gorm.io/gorm"
)

type transactionRepository struct {
	repository.CommonBehaviourRepository[entity.Transaction]
}

func NewTransactionRepository(db *gorm.DB) repository.TransactionRepository {
	return &transactionRepository{
		CommonBehaviourRepository: repository.NewCommonBehaviour[entity.Transaction](db),
	}
}
