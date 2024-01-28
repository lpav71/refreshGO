package models

import (
	"net/http"
	"strconv"
)

type Map struct {
	ID           int `gorm:"primaryKey"`
	ClubID       int
	UserID       int
	IDComp       int
	Zone         int
	Level        int
	PosX         int
	PosY         int
	StatusActive bool `gorm:"default:true"`
	IP           string
	MAC          string
	Ver          int
	Ping         int `gorm:"default:0"`
	MB           string
	CPU          string
	GPU          string
	RAM          string
	Disk         string
	Temp         string
	StatusNum    int `gorm:"default:0"`
	QRToLogin    string
	QRGen        int `gorm:"default:1"`
	ULogin       string
	UPass        string
	SOS          bool `gorm:"default:false"`
}

func (Map) TableName() string {
	return "map"
}

func (Map) SaveMapPosition(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clubId := r.Form.Get("club_id")
	comp := r.Form.Get("comp")
	posX := r.Form.Get("posx")
	posY := r.Form.Get("posy")

	posXInt, _ := strconv.Atoi(posX)
	posYInt, _ := strconv.Atoi(posY)

	// Создаем подзапрос
	subquery := Database.Model(&Map{}).
		Select("id").
		Where("id_comp = ? AND club_id = ?", comp, clubId).
		Limit(1)

	// Обновляем запись
	Database.Model(&Map{}).
		Where("id_comp = ? AND club_id = ? AND id = (?)", comp, clubId, subquery).
		Updates(Map{PosX: posXInt, PosY: posYInt})

}

func (Map) SaveNewPosition(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clubId := r.Form.Get("club_id")
	comp := r.Form.Get("id_comp")
	zone := r.Form.Get("zone")
	posX := r.Form.Get("posx")
	posY := r.Form.Get("posy")
	posXInt, _ := strconv.Atoi(posX)
	posYInt, _ := strconv.Atoi(posY)
	zoneInt, _ := strconv.Atoi(zone)
	compInt, _ := strconv.Atoi(comp)
	clubIdInt, _ := strconv.Atoi(clubId)

	map1 := Map{
		PosX:   posXInt,
		PosY:   posYInt,
		Zone:   zoneInt,
		IDComp: compInt,
		ClubID: clubIdInt,
		UserID: 440,
	}

	err := Database.Create(&map1)
	if err.Error != nil {
		id := 0
		idBytes := []byte(strconv.Itoa(id))
		w.Write(idBytes)
	} else {
		id := map1.ID
		idBytes := []byte(strconv.Itoa(id))

		w.Write(idBytes)
	}
}
