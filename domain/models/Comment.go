package models

import "time"

// Comment コメントの保管
type Comment struct {
	ID        uint32    `json:"id"`
	Text      string    `json:"text"`
	User      User      `json:"user"`
	UserID    uint32    `json:"user_id"`
	PostID    uint32    `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
