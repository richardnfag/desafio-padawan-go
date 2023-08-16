package services

import (
	"errors"
	"math"

	"github.com/richardnfag/desafio-padawan-go/internal/entities"
	"github.com/richardnfag/desafio-padawan-go/internal/ports"
)

type ConversionService interface {
	Convert(amount float64, fromCurrency, toCurrency string, rate float64) (*entities.Conversion, error)
}

type DefaultConversionService struct {
	conversionRepository ports.ConversionRepository
	currencyRepository   ports.CurrencyRepository
}

func NewConversionService(conversionRepository ports.ConversionRepository, currencyRepository ports.CurrencyRepository) *DefaultConversionService {
	return &DefaultConversionService{
		conversionRepository: conversionRepository,
		currencyRepository:   currencyRepository,
	}
}

func (s *DefaultConversionService) roundAmountValue(value float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(value*ratio) / ratio
}

func (s *DefaultConversionService) Convert(amount float64, fromCurrency, toCurrency string, rate float64) (*entities.Conversion, error) {

	convertedAmount := s.roundAmountValue(amount*rate, 3)

	fromCurrencyInfo, err := s.currencyRepository.GetCurrencyByCode(fromCurrency)
	if err != nil {
		return nil, err
	}

	toCurrencyInfo, err := s.currencyRepository.GetCurrencyByCode(toCurrency)
	if err != nil {
		return nil, err
	}

	if fromCurrencyInfo == nil || toCurrencyInfo == nil {
		return nil, errors.New("invalid currency code")
	}

	conversion := &entities.Conversion{
		Amount: amount,
		From:   *fromCurrencyInfo,
		To:     *toCurrencyInfo,
		Rate:   rate,
		Result: convertedAmount,
	}

	if err := s.conversionRepository.SaveConversion(conversion); err != nil {
		return nil, err
	}

	return conversion, nil
}
