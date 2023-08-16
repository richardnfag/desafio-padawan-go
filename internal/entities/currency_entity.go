package entities

import "gorm.io/gorm"

type Currency struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey"`
	Code   string `gorm:"not null; unique"`
	Symbol string
}
