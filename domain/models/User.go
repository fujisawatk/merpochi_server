package models

import (
	"merpochi_server/interfaces/security"
	"time"
)

// User ユーザー値の保管
type User struct {
	ID        uint32    `json:"id"`
	Nickname  string    `json:"nickname" validate:"required,min=1,max=20"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6,max=20"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeSave 保存前のパスワード処理
func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
