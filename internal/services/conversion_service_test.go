package services

import (
	"testing"

	"github.com/richardnfag/desafio-padawan-go/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockConversionRepository struct {
	mock.Mock
}

func NewMockConversionRepository() *MockConversionRepository {
	return &MockConversionRepository{
		Mock: mock.Mock{},
	}
}

func (m *MockConversionRepository) SaveConversion(conversion *entities.Conversion) error {
	return nil
}

type MockCurrencyRepository struct {
	mock.Mock
}

func NewMockCurrencyRepository() *MockCurrencyRepository {
	return &MockCurrencyRepository{
		Mock: mock.Mock{},
	}
}

func (m *MockCurrencyRepository) SaveCurrency(currency *entities.Currency) error {
	return nil
}

func (m *MockCurrencyRepository) SaveCurrencies(currencies *[]entities.Currency) error {
	return nil
}

func (m *MockCurrencyRepository) GetCurrencyByCode(code string) (*entities.Currency, error) {
	switch code {
	case "USD":
		return &entities.Currency{Code: "USD", Symbol: "$"}, nil
	case "BRL":
		return &entities.Currency{Code: "BRL", Symbol: "R$"}, nil
	case "EUR":
		return &entities.Currency{Code: "EUR", Symbol: "€"}, nil
	case "BTC":
		return &entities.Currency{Code: "BTC", Symbol: "₿"}, nil
	default:
		return nil, nil
	}
}

func TestConvert(t *testing.T) {
	conversionRepo := NewMockConversionRepository()
	currencyRepo := NewMockCurrencyRepository()

	service := NewConversionService(conversionRepo, currencyRepo)

	amount := 100.0
	fromCurrency := "USD"
	toCurrency := "EUR"
	rate := 0.85

	expectedFromCurrency := &entities.Currency{Code: "USD", Symbol: "$"}
	expectedToCurrency := &entities.Currency{Code: "EUR", Symbol: "€"}

	conversion, err := service.Convert(amount, fromCurrency, toCurrency, rate)

	assert.NoError(t, err)
	assert.NotNil(t, conversion)
	assert.Equal(t, amount, conversion.Amount)
	assert.Equal(t, *expectedFromCurrency, conversion.From)
	assert.Equal(t, *expectedToCurrency, conversion.To)
	assert.Equal(t, rate, conversion.Rate)

	currencyRepo.AssertExpectations(t)
	conversionRepo.AssertExpectations(t)
}

func TestConvertWithInvalidFromCurrency(t *testing.T) {
	conversionRepo := NewMockConversionRepository()
	currencyRepo := NewMockCurrencyRepository()

	service := NewConversionService(conversionRepo, currencyRepo)

	amount := 100.0
	fromCurrency := "UNK"
	toCurrency := "EUR"
	rate := 0.85

	_, err := service.Convert(amount, fromCurrency, toCurrency, rate)

	if assert.Error(t, err) {
		assert.Equal(t, "invalid currency code", err.Error())
	}
}

func TestConvertInvalidToCurrency(t *testing.T) {
	conversionRepo := NewMockConversionRepository()
	currencyRepo := NewMockCurrencyRepository()

	service := NewConversionService(conversionRepo, currencyRepo)

	amount := 100.0
	fromCurrency := "USD"
	toCurrency := "UNK"
	rate := 0.85

	_, err := service.Convert(amount, fromCurrency, toCurrency, rate)

	if assert.Error(t, err) {
		assert.Equal(t, "invalid currency code", err.Error())
	}
}
