package models

import (
	"encoding/json"
	"net/http"
)

func GetMapPosition(w http.ResponseWriter, r *http.Request) {
	// Разбор параметров запроса
	r.ParseForm()
	clubId := r.Form.Get("club_id")

	// Создание канала для передачи данных о зонах
	var zoneChan = make(chan []Zone, 1)
	var zone2Chan = make(chan []Zone, 1)
	var mapChan = make(chan []Map, 1)

	// Запуск горутины для получения
	go getZone(clubId, zone2Chan)
	go getMap(clubId, mapChan)
	go getMapZone(clubId, zoneChan)

	// Получение данных из каналов
	zones := <-zoneChan
	zone2 := <-zone2Chan
	map1 := <-mapChan

	outData := map[string]interface{}{
		"zone":    zone2,
		"map":     map1,
		"mapZone": zones,
	}

	// Преобразование данных в формат JSON
	zoneData, err := json.Marshal(outData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Отправка данных в формате JSON в ответ на запрос
	w.Header().Set("Content-Type", "application/json")
	w.Write(zoneData)
}

func getMapZone(clubId string, zoneChan chan []Zone) {
	// Получение данных о зонах из базы данных
	var zones []Zone
	err := Database.Table("zone").
		Select("zone.*").
		Joins("RIGHT JOIN map ON zone.club_id = map.club_id AND zone.num = map.zone").
		Where("map.club_id = ?", clubId).
		Order("id_comp").
		Find(&zones).Error

	if err != nil {
		// Обработка ошибки или отправка ошибки в канал
		zoneChan <- nil
		return
	}

	// Отправка данных о зонах в канал
	zoneChan <- zones
}

func getZone(clubId string, zone2Chan chan []Zone) {
	var zone []Zone
	err := Database.Where("club_id = ?", clubId).Find(&zone).Error
	if err != nil {
		// Обработка ошибки или отправка ошибки в канал
		zone2Chan <- nil
		return
	}

	// Отправка данных о зонах в канал
	zone2Chan <- zone
}

func getMap(clubId string, map1 chan []Map) {
	var maps []Map
	err := Database.Where("club_id = ?", clubId).Find(&maps).Error
	if err != nil {
		// Обработка ошибки или отправка ошибки в канал
		map1 <- nil
		return
	}

	// Отправка данных о зонах в канал
	map1 <- maps
}
