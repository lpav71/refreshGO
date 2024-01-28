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
		Select("task.id, club_id, admin_id, TO_CHAR(create_dt, 'MM.DD.YYYY HH:MI') as create_dt, employ, descript_admin, color").
		Joins("JOIN task_description ON task_description.id = task.status").
		Where("club_id = ? AND admin_id = ? AND status IN (?)", clubId, 7, []int{1, 2}).
		Order("create_dt").
		Find(&tasks).Error

	if err != nil {
		return
	}
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
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

	datePublicParsed, _ := time.Parse("2006-01-02", datePublic)
	daysInt, _ := strconv.Atoi(days)
	hoursInt, _ := strconv.Atoi(hours)

	endDT := datePublicParsed.AddDate(0, 0, daysInt).Add(time.Duration(hoursInt) * time.Hour)

	var task Task
	Database.First(&task, "id = ?", idTask)
	task.DescriptAdmin = &descriptAdmin
	task.CreateDt = &datePublicParsed
	task.EndDt = &endDT

	executorInt, _ := strconv.Atoi(executor)
	taskEmploy := executorInt
	task.Employ = &taskEmploy

	adminIDInt, _ := strconv.Atoi(adminID)
	taskAdminID := adminIDInt
	task.AdminID = &taskAdminID

	clubIDInt, _ := strconv.Atoi(clubID)
	taskClubID := clubIDInt
	task.ClubID = &taskClubID

	status := 1
	task.Status = &status

	err := Database.Save(&task)
	if err.Error != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		jsonData := []byte(`{"message": "Bad Request"}`)
		w.Write(jsonData)
	} else {
		jsonData, _ := json.Marshal(task)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
