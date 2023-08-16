package ports

import "github.com/richardnfag/desafio-padawan-go/internal/entities"

type CurrencyRepository interface {
	SaveCurrency(currency *entities.Currency) error
	GetCurrencyByCode(code string) (*entities.Currency, error)
}
