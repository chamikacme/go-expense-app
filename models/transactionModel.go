package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}
