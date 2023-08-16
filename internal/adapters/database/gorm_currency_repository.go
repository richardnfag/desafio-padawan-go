package database

import (
	"errors"

	"github.com/richardnfag/desafio-padawan-go/internal/entities"
	"gorm.io/gorm"
)

type GormCurrencyRepository struct {
	db *gorm.DB
}

func NewGormCurrencyRepository(db *gorm.DB) *GormCurrencyRepository {
	return &GormCurrencyRepository{db: db}
}

func (r *GormCurrencyRepository) SaveCurrency(currency *entities.Currency) error {
	return r.db.Create(currency).Error
}

func (r *GormCurrencyRepository) GetCurrencyByCode(code string) (*entities.Currency, error) {
	var currency entities.Currency

	err := r.db.Where("code = ?", code).First(&currency).Error

	if err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("currency code " + code + " not found")
		}

		return nil, err
	}

	return &currency, nil

}
