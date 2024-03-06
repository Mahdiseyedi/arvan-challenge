package entity

import "gorm.io/gorm"

type DiscountCode struct {
	gorm.Model

	Value string `json:"value"`
}

func (u DiscountCode) Table() string {
	return "DiscountCode"
}
