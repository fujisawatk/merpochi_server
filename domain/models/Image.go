package models

import (
	"bytes"
	"time"
)

// Image 画像情報の保管
type Image struct {
	ID        uint32        `json:"id"`
	Name      string        `json:"name"`
	UserID    uint32        `json:"user_id"`
	ShopID    uint32        `json:"shop_id"`
	Buf       *bytes.Buffer `json:"buf"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
