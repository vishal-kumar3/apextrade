package repository

import (
	"github.com/apextrade/config"
	"github.com/apextrade/internal/models"
	"gorm.io/gorm"
)

type PostgresOrderRepo struct {
	db *gorm.DB
}

func NewPostgresOrderRepo(d *config.DB) *PostgresOrderRepo {
	return &PostgresOrderRepo{db: d.DB}
}

func (r *PostgresOrderRepo) CreateOrder(o *models.Order) error {
	if err := o.Validate(); err != nil {
		return err
	}
	return r.db.Create(o).Error
}

func (r *PostgresOrderRepo) GetByID(id uint) (models.Order, bool) {
	var o models.Order
	result := r.db.First(&o, id)
	return o, result.Error == nil && result.RowsAffected > 0
}

func (r *PostgresOrderRepo) GetBySymbol(symbol string) []models.Order {
	var orders []models.Order
	r.db.Where("symbol = ?", symbol).Find(&orders)
	return orders
}

func (r *PostgresOrderRepo) UpdateStatus(id uint, status models.OrderStatus) bool {
	result := r.db.Model(&models.Order{}).Where("id = ?", id).Update("status", status)
	return result.Error == nil && result.RowsAffected > 0
}

func (r *PostgresOrderRepo) DeleteById(id uint) bool {
	result := r.db.Delete(&models.Order{}, id)
	return result.RowsAffected > 0
}
