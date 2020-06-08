package models

import "time"

// Shop 店舗IDの保管
type Shop struct {
	ID        uint32    `json:"id"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
