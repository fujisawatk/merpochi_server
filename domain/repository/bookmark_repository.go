package repository

import "merpochi_server/domain/models"

// BookmarkRepository bookmarkPersistenceの抽象依存
type BookmarkRepository interface {
	Save(*models.Bookmark) (*models.Bookmark, error)
}
