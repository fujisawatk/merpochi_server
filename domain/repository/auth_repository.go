package repository

import "merpochi_server/domain/models"

// AuthRepository userPersistenceの抽象依存
type AuthRepository interface {
	SignIn(string, string) (models.User, string, error)
	FindCurrentUser(uint32) (models.User, string, error)
}
