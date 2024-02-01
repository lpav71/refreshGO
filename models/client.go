package models

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type ClientTable struct {
	ID     int    `gorm:"primary_key"`
	ClubID int    `gorm:"column:club_id"`
	Login  string `gorm:"not null"`
	//Password     string // Если нужно по умолчанию NULL поле можно не указывать
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

// TableName Указание имени таблицы для модели Client
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
func (ClientTable) GetClients(w http.ResponseWriter, r *http.Request) {
	clubID := r.FormValue("club_id")
	var clients []Client
	Database.Where("club_id = ?", clubID).Find(&clients)
	jsonData, _ := json.Marshal(clients)
	w.Write(jsonData)
}
func (ClientTable) GetByLogin(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	var client ClientTable
	Database.Where("login = ?", login).First(&client)
	jsonData, _ := json.Marshal(client)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func (ClientTable) UsersFind(w http.ResponseWriter, r *http.Request) {
	clubID := r.FormValue("club_id")
	searchUser := r.FormValue("searchUser")
	clubIDInt, _ := strconv.ParseInt(clubID, 10, 32)
	var clients []ClientTable
	Database.Where("login ILIKE ?", "%"+searchUser+"%").Where("club_id = ?", clubIDInt).Find(&clients)
	jsonData, _ := json.Marshal(clients)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func (ClientTable) AgeFilter(w http.ResponseWriter, r *http.Request) {
	age := r.FormValue("age")
	var clients []Client

	switch age {
	case "0-12":
		Database.Where("age(bday) BETWEEN '0 years' AND '12 years'").Find(&clients)
	case "12-16":
		Database.Where("age(bday) BETWEEN '12 years' AND '16 years'").Find(&clients)
	case "16-18":
		Database.Where("age(bday) BETWEEN '16 years' AND '18 years'").Find(&clients)
	case "18-24":
		Database.Where("age(bday) BETWEEN '18 years' AND '24 years'").Find(&clients)
	case "25-29":
		Database.Where("age(bday) BETWEEN '24 years' AND '29 years'").Find(&clients)
	case "30-999":
		Database.Where("age(bday) BETWEEN '30 years' AND '999 years'").Find(&clients)
	}
	jsonData, _ := json.Marshal(clients)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
