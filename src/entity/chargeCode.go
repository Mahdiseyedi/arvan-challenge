package entity

import "gorm.io/gorm"

type ChargeCode struct {
	gorm.Model

	Value string `json:"value"`
}

func (u ChargeCode) Table() string {
	return "ChargeCode"
}
