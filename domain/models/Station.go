package models

import "time"

// Station 駅名の保管
type Station struct {
	ID           uint32    `json:"id"`
	StationName  string    `json:"station_name"`
	StationNameK string    `json:"station_name_k"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
