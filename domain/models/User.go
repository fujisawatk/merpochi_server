package models

import "time"

// User ユーザー値の保管
type User struct {
	ID        uint32
	Nickname  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
