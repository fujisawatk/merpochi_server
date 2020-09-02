package repository

import "merpochi_server/domain/models"

// StationRepository stationPersistenceの抽象依存
type StationRepository interface {
	SearchKanaWord(string) ([]models.Station, error)
	SearchKanjiWord(string) ([]models.Station, error)
}
