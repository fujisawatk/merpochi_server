package usecase

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
)

// StationUsecase Stationに対するUsecaseのインターフェイス
type StationUsecase interface {
	SearchStations(string) ([]models.Station, error)
}

type stationUsecase struct {
	stationRepository repository.StationRepository
}

// NewStationUsecase Stationデータに関するUsecaseを生成
func NewStationUsecase(sr repository.StationRepository) StationUsecase {
	return &stationUsecase{
		stationRepository: sr,
	}
}

func (su stationUsecase) SearchStations(word string) ([]models.Station, error) {
	stations, err := su.stationRepository.SearchKanjiWord(word)
	if err != nil {
		return nil, err
	}
	return stations, nil
}
