package repository

import "merpochi_server/domain/models"

// UserRepository userPersistenceの抽象依存
type UserRepository interface {
	Save(models.User) (models.User, error)
}
