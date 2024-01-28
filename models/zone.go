package models

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Zone struct {
	ID     uint `gorm:"primaryKey"`
	ClubID uint
	Num    uint
	Name   string
}

func (Zone) TableName() string {
	return "zone"
}

func (Zone) AddZone(w http.ResponseWriter, r *http.Request) {
	clubId := r.FormValue("club_id")
	zoneName := r.FormValue("zone_name")
	if clubId != "" && zoneName != "" {
		var maximum int
		Database.Model(&Zone{}).Select("MAX(num) as max").Where("club_id = ?", clubId).Scan(&maximum)
		clubIdInt, _ := strconv.Atoi(clubId)
		zone := Zone{
			Num:    uint(maximum + 1),
			ClubID: uint(clubIdInt),
			Name:   zoneName,
		}
		err := Database.Create(&zone)

		if err.Error != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			jsonData := []byte(`{"message": "Bad Request"}`)
			w.Write(jsonData)
		} else {
			jsonData, _ := json.Marshal(zone)

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}
}
