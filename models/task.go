package models

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	ID               int
	ClubID           *int       `gorm:"column:club_id"`
	AdminID          *int       `gorm:"column:admin_id"`
	CreateDt         *time.Time `gorm:"column:create_dt"`
	Employ           *int       `gorm:"column:employ"`
	TaskID           *int       `gorm:"column:task_id"`
	EndDt            *time.Time `gorm:"column:end_dt"`
	DescriptAdmin    *string    `gorm:"column:descript_admin"`
	DescriptExecutor *string    `gorm:"column:descript_executor"`
	Status           *int       `gorm:"column:status"`
	Visible          bool       `gorm:"column:visible;default:true"`
}

type Task1 struct {
	ID               *int
	ClubID           *int
	AdminID          *int   `gorm:"column:admin_id"`
	CreateDT         string `gorm:"column:create_dt"`
	Employ           *int
	TaskID           *int `gorm:"column:end_dt"`
	EndDt            *time.Time
	DescriptAdmin    *string `gorm:"column:descript_admin"`
	DescriptExecutor *string `gorm:"column:descript_executor"`
	Status           *int
	Visible          bool
	Color            *string
}

func (Task) TableName() string {
	return "task"
}

func (Task) Tasks(w http.ResponseWriter, r *http.Request) {
	clubId := r.FormValue("club_id")
	var tasks []Task1

	err := Database.Table("task").
		Select("task.id, club_id, admin_id, TO_CHAR(create_dt, 'DD.MM.YYYY HH24:MI') as create_dt, employ, descript_admin, color").
		Joins("JOIN task_description ON task_description.id = task.status").
		Where("club_id = ? AND admin_id = ? AND status IN (?)", clubId, 7, []int{1, 2}).
		Order("create_dt").
		Find(&tasks).Error

	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, tasks)
}

func (Task) SaveEditModal(w http.ResponseWriter, r *http.Request) {
	idTask := r.FormValue("id_task")
	descriptAdmin := r.FormValue("descript_admin")
	datePublic := r.FormValue("date_public")
	days := r.FormValue("days")
	hours := r.FormValue("hours")
	executor := r.FormValue("executor")
	clubID := r.FormValue("club_id")
	adminID := r.FormValue("admin_id")

	datePublicParsed, err := time.Parse("02.01.2006 15:04", datePublic)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	endDT := datePublicParsed.AddDate(0, 0, StrToInt(days)).Add(time.Duration(StrToInt(hours)) * time.Hour)

	var task Task
	if err := Database.First(&task, idTask).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	task.DescriptAdmin = &descriptAdmin
	task.CreateDt = &datePublicParsed
	task.EndDt = &endDT
	task.Employ = intPtr(StrToInt(executor))
	task.AdminID = intPtr(StrToInt(adminID))
	task.ClubID = intPtr(StrToInt(clubID))

	if err := Database.Save(&task).Error; err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	respondWithJSON(w, task)
}

func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func intPtr(i int) *int {
	return &i
}

func respondWithJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if jsonData, err := json.Marshal(data); err == nil {
		w.Write(jsonData)
	} else {
		http.Error(w, "Failed to convert to JSON", http.StatusInternalServerError)
	}
}
