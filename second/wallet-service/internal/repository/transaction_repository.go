package repository

import "github.com/arvan-challenge/wallet/internal/entity"

type TransactionRepository interface {
	CreateTransaction(transaction *entity.Transaction) error
	GetTransactionByID(transactionID uint) (*entity.Transaction, error)
	DeleteTransaction(transactionID uint) error
	GetUserTransactions(userID uint) ([]*entity.Transaction, error)
}
