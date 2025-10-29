package repository

import (
	"math/rand"
	"time"

	"github.com/apextrade/internal/models"
)

type InMemoryStockRepo struct {
	stocks map[string]models.Stock
}

func NewInMemoryStockRepo() *InMemoryStockRepo {
	// seed data
	return &InMemoryStockRepo{
		stocks: map[string]models.Stock{
			"AAPL": {Symbol: "AAPL", Price: 150.25, Volume: 1000000, UpdatedAt: time.Now()},
			"GOOG": {Symbol: "GOOG", Price: 2800.50, Volume: 500000, UpdatedAt: time.Now()},
		},
	}
}

func (r *InMemoryStockRepo) GetBySymbol(symbol string) (models.Stock, bool) {
	if stock, ok := r.stocks[symbol]; ok {
		return stock, true
	}
	return models.Stock{}, false
}

func (r *InMemoryStockRepo) GetAll() []models.Stock {
	stocks := make([]models.Stock, 0, len(r.stocks))
	for _, s := range r.stocks {
		stocks = append(stocks, s)
	}
	return stocks
}

func (r *InMemoryStockRepo) CreateOrUpdate(s *models.Stock) {
	s.UpdatedAt = time.Now()
	if s.Price == 0 {
		s.Price = 100 + rand.Float64()*200
	}
	r.stocks[s.Symbol] = *s
}
