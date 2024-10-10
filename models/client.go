package models

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type ClientTable struct {
	ID           int     `gorm:"primary_key"`
	ClubID       *int    `gorm:"column:club_id"`
	Login        string  `gorm:"not null"`
	Password     *string // Если нужно по умолчанию NULL поле можно не указывать
	Phone        string  `gorm:"not null"`
	Email        string  `gorm:"not null;unique_index:clients_email"`
	Icon         *string
	Amount       *float64 `gorm:"default:0"`
	Bonus        *float64 `gorm:"default:0"`
	TotalTime    *int     `gorm:"default:0"`
	FullName     *string
	StatusActive *bool      `gorm:"column:status_active"`
	TelegramID   *string    `gorm:"column:telegram_id"`
	VKID         *string    `gorm:"column:vk_id"`
	RegDate      *time.Time `gorm:"default:now()"`
	BDay         *time.Time `gorm:"column:bday"`
	Verify       *bool      `gorm:"default:false"`
	VerifyDt     *time.Time `gorm:"column:verify_dt"`
	MiddleName   *string    `gorm:"column:middle_name"`
	Surname      *string    `gorm:"column:surname"`
	Name         *string    `gorm:"column:name"`
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
func (ClientTable) ExportClients(w http.ResponseWriter, r *http.Request) {
	clubId := r.FormValue("club_id")
	var clients []ClientTable
	var fileName string
	Database.Where("club_id = ?", clubId).Find(&clients)
	fileName = runExportClients(clients)
	w.Write([]byte(fileName))
}
func (ClientTable) ImportClients(w http.ResponseWriter, r *http.Request) {
	fileName, _, _ := r.FormFile("filename")
	var dataFromFile [][]string
	dataFromFile = readCSVFile(fileName)
	insertNewRecords(dataFromFile)
}

func readCSVFile(fileName multipart.File) [][]string {

	defer fileName.Close()

	reader := csv.NewReader(fileName)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var clientsFromFile [][]string
	var ids []int
	for _, line := range records {
		clientsFromFile = append(clientsFromFile, line)
		id, _ := strconv.Atoi(line[0])
		ids = append(ids, id)
	}
	return clientsFromFile
}
func getValueString[T any](ptr *T, formatFunc func(T) string) string {
	if ptr == nil {
		return ""
	}
	return formatFunc(*ptr)
}
func runExportClients(clients []ClientTable) string {
	date := time.Now().Format("2006-01-02 15:04:05")
	fileName := fmt.Sprintf("%x.csv", md5.Sum([]byte(date)))

	file, err := os.Create("views/files/" + fileName)
	if err != nil {
		return ""
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	for _, user := range clients {
		err := writer.Write([]string{
			strconv.Itoa(user.ID),
			getValueString(user.ClubID, func(v int) string { return strconv.Itoa(v) }),
			user.Login,
			getValueString(user.Password, func(v string) string { return v }),
			user.Phone,
			user.Email,
			getValueString(user.Icon, func(v string) string { return v }),
			getValueString(user.Amount, func(v float64) string { return strconv.FormatFloat(v, 'f', -1, 64) }),
			getValueString(user.Bonus, func(v float64) string { return strconv.FormatFloat(v, 'f', -1, 64) }),
			getValueString(user.TotalTime, func(v int) string { return strconv.Itoa(v) }),
			getValueString(user.FullName, func(v string) string { return v }),
			getValueString(user.StatusActive, func(v bool) string { return strconv.FormatBool(v) }),
			getValueString(user.TelegramID, func(v string) string { return v }),
			getValueString(user.VKID, func(v string) string { return v }),
			getValueString(user.RegDate, func(v time.Time) string { return v.Format("2006-01-02") }),
			getValueString(user.BDay, func(v time.Time) string { return v.Format("2006-01-02") }),
			getValueString(user.Verify, func(v bool) string { return strconv.FormatBool(v) }),
			getValueString(user.VerifyDt, func(v time.Time) string { return v.Format("2006-01-02 15:04:05") }),
			getValueString(user.MiddleName, func(v string) string { return v }),
			getValueString(user.Surname, func(v string) string { return v }),
			getValueString(user.Name, func(v string) string { return v }),
			strconv.Itoa(user.GroupID),
			strconv.FormatFloat(user.GroupAmount, 'f', -1, 64),
			getValueString(user.GroupCreate, func(v time.Time) string { return v.Format("2006-01-02 15:04:05") }),
		})

		if err != nil {
			return ""
		}
	}
	return fileName
}
func insertNewRecords(clientsFromFile [][]string) {
	for _, clientFromFile := range clientsFromFile {
		var client ClientTable
		if err := Database.Where("login = ?", clientFromFile[2]).First(&client).Error; err != nil && strings.Contains(err.Error(), "record not found") {
			groupCreate := parseOptionalDate(clientFromFile[24])
			clubID := parseInt(clientFromFile[1])
			amount := parseFloat(clientFromFile[7])
			bonus := parseFloat(clientFromFile[8])
			totalTime := parseInt(clientFromFile[9])
			statusActive := parseBool(clientFromFile[11])
			regDate := parseOptionalDate(clientFromFile[14])
			verify := parseOptionalBool(clientFromFile[16])
			groupID := parseInt(clientFromFile[21])
			groupAmount := parseFloat(clientFromFile[22])

			client := ClientTable{
				ClubID:       clubID,
				Login:        clientFromFile[2],
				Password:     &clientFromFile[3],
				Phone:        clientFromFile[4],
				Email:        clientFromFile[5],
				Icon:         &clientFromFile[6],
				Amount:       amount,
				Bonus:        bonus,
				TotalTime:    totalTime,
				FullName:     &clientFromFile[10],
				StatusActive: &statusActive,
				TelegramID:   &clientFromFile[12],
				VKID:         &clientFromFile[13],
				RegDate:      regDate,
				BDay:         parseDateTime(clientFromFile[15]),
				Verify:       verify,
				VerifyDt:     parseDateTime(clientFromFile[17]),
				Name:         &clientFromFile[18],
				Surname:      &clientFromFile[19],
				MiddleName:   &clientFromFile[20],
				GroupID:      *groupID,
				GroupAmount:  *groupAmount,
				GroupCreate:  groupCreate,
			}
			Database.Create(&client)
		}
	}
}

func parseOptionalDate(dateStr string) *time.Time {
	if dateStr == "" {
		return nil
	}
	timeParsed, _ := time.Parse("2006-01-02", dateStr)
	return &timeParsed
}
func parseInt(str string) *int {
	if val, err := strconv.Atoi(str); err == nil {
		return &val
	}
	return nil
}
func parseFloat(str string) *float64 {
	if val, err := strconv.ParseFloat(str, 64); err == nil {
		return &val
	}
	return nil
}
func parseBool(str string) bool {
	result, _ := strconv.ParseBool(str)
	return result
}
func parseOptionalBool(str string) *bool {
	if str == "" {
		return nil
	}
	result := parseBool(str)
	return &result
}
func parseDateTime(dateTimeStr string) *time.Time {
	if dateTimeStr == "" {
		return nil
	}
	parsedDateTime, _ := time.Parse("2006-01-02 15:04:05", dateTimeStr)
	return &parsedDateTime
}
