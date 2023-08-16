package entities

import "gorm.io/gorm"

type Conversion struct {
	gorm.Model
	ID             uint `gorm:"primaryKey"`
	Amount         float64
	FromCurrencyID uint
	From           Currency `gorm:"foreignKey:FromCurrencyID;references:ID"`
	ToCurrencyID   uint
	To             Currency `gorm:"foreignKey:ToCurrencyID;references:ID"`
	Rate           float64
	Result         float64
}
