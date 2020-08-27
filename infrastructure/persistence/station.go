package persistence

import (
	"merpochi_server/domain/models"
	"merpochi_server/domain/repository"

	"merpochi_server/util/channels"

	"github.com/jinzhu/gorm"
)

type stationPersistence struct {
	db *gorm.DB
}

// NewStationPersistence stationPersistence構造体の宣言
func NewStationPersistence(db *gorm.DB) repository.StationRepository {
	return &stationPersistence{db}
}

func (sp *stationPersistence) SearchKanaWord(word string) ([]models.Station, error) {
	var err error

	stations := []models.Station{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = sp.db.Debug().Model(&models.Station{}).Where("convert(station_name_k using utf8) collate utf8_unicode_ci LIKE ?", "%"+word+"%").Find(&stations).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return stations, nil
	}
	return nil, err
}

func (sp *stationPersistence) SearchKanjiWord(word string) ([]models.Station, error) {
	var err error

	stations := []models.Station{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err = sp.db.Debug().Model(&models.Station{}).Where("station_name LIKE ?", "%"+word+"%").Find(&stations).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return stations, nil
	}
	return nil, err
}
