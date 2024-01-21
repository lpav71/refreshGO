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
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

func (Zone) TableName() string {
	return "zone"
}

func (Zone) AddZone(w http.ResponseWriter, r *http.Request) {
	clubId := r.FormValue("club_id")
	zoneName := r.FormValue("zone_name")
	if clubId != "" && zoneName != "" {
		var maximum int
		Database.Model(&Zone{}).Select("MAX(num) as max").Where("club_id = ?", 1).Scan(&maximum)
		clubIdInt, _ := strconv.Atoi(clubId)
		zone := Zone{
			Num:    uint(maximum + 1),
			ClubID: uint(clubIdInt),
			Name:   zoneName,
		}
		Database.Create(&zone)
		jsonData, _ := json.Marshal(zone)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
