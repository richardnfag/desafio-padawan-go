package ports

import "github.com/richardnfag/desafio-padawan-go/internal/entities"

type ConversionRepository interface {
	SaveConversion(conversion *entities.Conversion) error
}
