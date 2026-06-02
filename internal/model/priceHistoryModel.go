package model

import "time"

type PriceHistoryDomain struct {
	ID         string    `json:"id"`
	ProductID  string    `json:"product_id"`
	Price      float64   `json:"price"`
	CapturedAt time.Time `json:"captured_at"`
}
