package models

import "time"

// Station 駅名の保管
type Station struct {
	ID           uint32    `json:"id"`
	StationName  string    `json:"station_name"`
	StationNameK string    `json:"station_name_k"`
	Prefecture   string    `json:"prefecture"`
	LineOne      string    `json:"line_one"`
	LineTwo      string    `json:"line_two"`
	CreatedAt    time.Time `json:"created_at"`
}
