package repository

import (
	"math/rand"
	"time"

	"github.com/apextrade/config"
	"github.com/apextrade/internal/models"
	"gorm.io/gorm"
)

type PostgresStockRepo struct {
	db *gorm.DB
}

func NewPostgresStockRepo(d *config.DB) *PostgresStockRepo {
	return &PostgresStockRepo{db: d.DB}
}

func (r *PostgresStockRepo) GetBySymbol(symbol string) (models.Stock, bool) {
	var stock models.Stock
	result := r.db.Where("symbol = ?", symbol).First(&stock)
	return stock, result.Error == nil && result.RowsAffected > 0
}

func (r *PostgresStockRepo) GetAll() []models.Stock {
	var stocks []models.Stock
	r.db.Find(&stocks)
	return stocks
}

func (r *PostgresStockRepo) CreateOrUpdate(s *models.Stock) error {
	if err := s.Validate(); err != nil {
		return err
	}
	s.UpdatedAt = time.Now()
	if s.Price == 0 {
		s.Price = 100 + rand.Float64()*200
	}
	return r.db.Save(s).Error
}

func (r *PostgresStockRepo) Delete(symbol string) bool {
	result := r.db.Where("symbol = ?", symbol).Delete(&models.Stock{})
	return result.RowsAffected > 0
}
