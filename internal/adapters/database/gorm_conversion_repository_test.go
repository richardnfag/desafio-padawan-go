package database

import (
	"testing"

	"github.com/richardnfag/desafio-padawan-go/internal/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGormConversionRepositorySaveConversion(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&entities.Conversion{}, &entities.Currency{})
	repo := NewGormConversionRepository(db)

	fromCurrency := &entities.Currency{
		Code:   "USD",
		Symbol: "$",
	}

	toCurrency := &entities.Currency{
		Code:   "EUR",
		Symbol: "€",
	}

	conversion := &entities.Conversion{
		Amount: 100,
		From:   *fromCurrency,
		To:     *toCurrency,
		Rate:   0.85,
	}

	err := repo.SaveConversion(conversion)

	assert.NoError(t, err)
	assert.NotZero(t, conversion.ID)
}
