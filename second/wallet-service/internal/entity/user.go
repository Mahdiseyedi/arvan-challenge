package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	PhoneNumber  string        `json:"phone_number" gorm:"unique"`
	Balance      float64       `json:"balance"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:UserID"`
}
