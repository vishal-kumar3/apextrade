package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Order struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Symbol    string    `json:"symbol" gorm:"type:varchar(10);not null" validate:"required,uppercase,len=1..10"`
	Side      string    `json:"side" gorm:"type:varchar(3);not null" validate:"required,oneof=buy sell"`
	Shares    int       `json:"shares" gorm:"not null" validate:"required,gt=0"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2);not null" validate:"required,gt=0"`
	Status    string    `json:"status" gorm:"type:varchar(20);default:'pending'" validate:"required,oneof=pending filled cancelled"`
	UserID    uint      `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (o *Order) Validate() error {
	v := validator.New()
	return v.Struct(o)
}
