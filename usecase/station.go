package usecase

import (
	"errors"
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"
	"merpochi_server/usecase/validations"
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
	// バリデーションで処理を分技
	format, err := validations.StationSearchValidate(word)
	switch format {
	case "HiraganaOrKatakana":
		stations, err := su.stationRepository.SearchKanaWord(word)
		if err != nil {
			return nil, err
		}
		return stations, nil
	case "AndKanji":
		stations, err := su.stationRepository.SearchKanjiWord(word)
		if err != nil {
			return nil, err
		}
		return stations, nil
	case "unknown":
		return []models.Station{}, err
	}
	return []models.Station{}, errors.New("予期しないエラーが発生しました")
}
