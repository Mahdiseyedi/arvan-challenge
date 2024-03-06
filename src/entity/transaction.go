package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	User   User    `json:"-"`
	Amount float64 `json:"amount" gorm:"-"`
}

func (u Transaction) Table() string {
	return "Transaction"
}
