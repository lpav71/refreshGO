package models

import (
	"encoding/json"
	"net/http"
	"sync"
)

type ResponseData struct {
	Zone    []Zone `json:"zone"`
	Map     []Map  `json:"map"`
	MapZone []Zone `json:"mapZone"`
}

func GetMapPosition(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(3)

	// Разбор параметров запроса
	r.ParseForm()
	clubId := r.Form.Get("club_id")

	// Структура для сбора данных
	var data ResponseData

	// Запуск горутин
	go func() {
		defer wg.Done()
		data.Zone = getZone(clubId)
	}()
	go func() {
		defer wg.Done()
		data.Map = getMap(clubId)
	}()
	go func() {
		defer wg.Done()
		data.MapZone = getMapZone(clubId)
	}()

	wg.Wait()

	// Преобразование данных в формат JSON
	zoneData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Отправка данных в формате JSON в ответ на запрос
	w.Header().Set("Content-Type", "application/json")
	w.Write(zoneData)
}

func getMapZone(clubId string) []Zone {
	var zones []Zone
	err := Database.Table("zone").
		Select("zone.*").
		Joins("RIGHT JOIN map ON zone.club_id = map.club_id AND zone.num = map.zone").
		Where("map.club_id = ?", clubId).
		Order("id_comp").
		Find(&zones).Error

	if err != nil {
		return nil // Или обработка ошибки
	}
	return zones
}

func getZone(clubId string) []Zone {
	var zone []Zone
	err := Database.Where("club_id = ?", clubId).Find(&zone).Error
	if err != nil {
		return nil // Или обработка ошибки
	}
	return zone
}

func getMap(clubId string) []Map {
	var maps []Map
	err := Database.Where("club_id = ?", clubId).Find(&maps).Error
	if err != nil {
		return nil // Или обработка ошибки
	}
	return maps
}
