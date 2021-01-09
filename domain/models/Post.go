package models

import "time"

// Post 投稿の保管
type Post struct {
	ID        uint32    `json:"id"`
	Text      string    `json:"text"`
	Rating    uint32    `json:"rating"`
	UserID    uint32    `json:"user_id"`
	ShopID    uint32    `json:"shop_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
