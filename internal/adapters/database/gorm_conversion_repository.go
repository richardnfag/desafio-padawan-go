package database

import (
	"github.com/richardnfag/desafio-padawan-go/internal/entities"
	"gorm.io/gorm"
)

type GormConversionRepository struct {
	db *gorm.DB
}

func NewGormConversionRepository(db *gorm.DB) *GormConversionRepository {
	return &GormConversionRepository{db: db}
}

func (r *GormConversionRepository) SaveConversion(conversion *entities.Conversion) error {
	return r.db.Create(conversion).Error
}
