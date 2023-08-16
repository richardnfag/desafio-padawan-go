package database

import (
	"testing"

	"github.com/richardnfag/desafio-padawan-go/internal/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func databaseSetup() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&entities.Currency{})
	return db
}

func TestGormCurrencyRepositorySaveCurrency(t *testing.T) {
	db := databaseSetup()
	repo := NewGormCurrencyRepository(db)

	currency := &entities.Currency{
		Code:   "USD",
		Symbol: "$",
	}

	err := repo.SaveCurrency(currency)

	assert.NoError(t, err)
	assert.NotZero(t, currency.ID)
}

func TestGormCurrencyRepositoryGetCurrencyByCode(t *testing.T) {
	db := databaseSetup()
	repo := NewGormCurrencyRepository(db)

	currencySample := &entities.Currency{
		Code:   "USD",
		Symbol: "$",
	}

	db.Create(currencySample)

	currencyRetrieved, err := repo.GetCurrencyByCode(currencySample.Code)

	assert.NoError(t, err)
	assert.NotZero(t, currencyRetrieved.ID)
	assert.Equal(t, currencySample.Code, currencyRetrieved.Code)
	assert.Equal(t, currencySample.Symbol, currencyRetrieved.Symbol)
}

func TestGormCurrencyRepositoryGetCurrencyByCodeNonExistent(t *testing.T) {
	db := databaseSetup()
	repo := NewGormCurrencyRepository(db)

	currencySampleCode := "USD"

	_, err := repo.GetCurrencyByCode(currencySampleCode)

	if assert.Error(t, err) {
		assert.Equal(t, "currency code USD not found", err.Error())
	}

}
