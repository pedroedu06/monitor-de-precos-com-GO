package model

import "time"

type NotificationDomain struct {
	ID            string    `json:"id"`
	ProductID     string    `json:"product_id"`
	PreviousPrice float64   `json:"previous_price"`
	CurrentPrice  float64   `json:"current_price"`
	SentAt        time.Time `json:"sent_at"`
}
