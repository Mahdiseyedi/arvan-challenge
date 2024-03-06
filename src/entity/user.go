package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model

	PhoneNumber string  `json:"phone_number"`
	Balance     float64 `json:"balance"`
}

func (u User) Table() string {
	return "User"
}
