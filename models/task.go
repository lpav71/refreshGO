package models

import (
	"encoding/json"
	"net/http"
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
	AdminID          *int
	CreateDT         string
	Employ           *int
	TaskID           *int
	EndDt            *time.Time
	DescriptAdmin    *string
	DescriptExecutor *string
	Status           *int
	Visible          bool
	Color            *string
}

func (Task) TableName() string {
	return "task"
}

func (Task) Tasks(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	//clubId := r.Form.Get("club_id")
	var tasks []Task1
	err = Database.Table("task").
		Select("task.id, club_id, admin_id, TO_CHAR(create_dt, 'MM.DD.YYYY HH:MI') as create_dt, employ, descript_admin, color").
		Joins("JOIN task_description ON task_description.id = task.status").
		Where("club_id = ? AND admin_id = ? AND status IN (?)", 1, 7, []int{1, 2}).
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
