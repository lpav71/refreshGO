package models

import (
	"encoding/json"
	"net/http"
	"time"
)

type ClientTable struct {
	ID           int    `gorm:"primary_key"`
	ClubID       int    `gorm:"column:club_id"`
	Login        string `gorm:"not null"`
	Password     string // Если нужно по умолчанию NULL поле можно не указывать
	Phone        string `gorm:"not null"`
	Email        string `gorm:"not null;unique_index:clients_email"`
	Icon         string
	Amount       float64 `gorm:"default:0"`
	Bonus        float64 `gorm:"default:0"`
	TotalTime    int     `gorm:"default:0"`
	FullName     string
	StatusActive *bool      `gorm:"column:status_active"`
	TelegramID   string     `gorm:"column:telegram_id"`
	VKID         string     `gorm:"column:vk_id"`
	RegDate      time.Time  `gorm:"default:now()"`
	BDay         *time.Time `gorm:"column:bday"`
	Verify       *bool      `gorm:"default:false"`
	VerifyDt     *time.Time `gorm:"column:verify_dt"`
	MiddleName   string     `gorm:"column:middle_name"`
	Surname      string     `gorm:"column:surname"`
	Name         string     `gorm:"column:name"`
	GroupID      int        `gorm:"not null;default:0"`
	GroupAmount  float64    `gorm:"not null;default:0"`
	GroupCreate  *time.Time `gorm:"column:group_create"`
}

// Указание имени таблицы для модели Client
func (ClientTable) TableName() string {
	return "clients"
}

type GetLoginJson struct {
	Login string
}

func (ClientTable) GetLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userId := r.Form.Get("user_id")
	if userId != "0" {
		var client ClientTable
		Database.Table("clients").Select("login").Where("id = ?", userId).Find(&client)
		data := GetLoginJson{
			Login: client.Login,
		}
		jsonData, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		data := GetLoginJson{
			Login: "",
		}
		jsonData, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
