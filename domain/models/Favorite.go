package models

import "time"

// Favorite お気に入り値の保管
type Favorite struct {
	ID        uint32    `json:"id"`
	User      User      `json:"user"`
	UserID    uint32    `json:"user_id"`
	ShopID    uint32    `json:"shop_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
