package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID               uint `gorm:"primaryKey"`
	Name             string
	Email            string
	EmailVerifiedAt  *time.Time
	Password         string
	Role             string `gorm:"default:user"`
	RememberToken    string
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	Surname          string
	NameAuth         string
	MiddleName       string
	Birthday         *time.Time
	Phone            string
	Icon             string
	Sex              string
	DateOfEmployment string
	ClubID           uint
	PositionID       uint
	GroupID          uint
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var user []User
	err := Database.Find(&user).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func (User) Users(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clubId := r.Form.Get("club_id")
	var user []User
	Database.Table("users").Where("club_id = ?", clubId).Find(&user)
	jsonData, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
