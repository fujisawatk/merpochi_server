package models

import "time"

// Comment コメントの保管
type Comment struct {
	ID        uint32    `json:"id"`
	Text      string    `json:"text"`
	ShopID    uint32    `json:"shop_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
