package models

import "time"

// Shop 店舗IDの保管
type Shop struct {
	ID        uint32    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Opentime  string    `json:"opentime"`
	Budget    uint32    `json:"budget"`
	Img       string    `json:"img"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
