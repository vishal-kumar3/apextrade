package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Stock struct {
	Symbol    string    `json:"symbol" gorm:"primaryKey;unique;type:varchar(10);not null" validate:"required,uppercase,min=1,max=10"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2);not null" validate:"gt=0"`
	Volume    int       `json:"volume" gorm:"default:0" validate:"min=0"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type Holding struct {
	Symbol string  `json:"symbol"`
	Shares int     `json:"shares"`
	AvgBuy float64 `json:"avg_buy"`
}

func (s *Stock) Validate() error {
	v := validator.New()
	return v.Struct(s)
}
