package models

import "time"

type Stock struct {
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price"`
	Volume    int       `json:"volume"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Holding struct {
	Symbol string  `json:"symbol"`
	Shares int     `json:"shares"`
	AvgBuy float64 `json:"avg_buy"`
}
