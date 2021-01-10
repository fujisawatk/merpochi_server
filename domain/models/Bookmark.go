package models

import "time"

// Bookmark お気に入り値の保管
type Bookmark struct {
	ID        uint32    `json:"id"`
	UserID    uint32    `json:"user_id"`
	ShopID    uint32    `json:"shop_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
