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
	RegDate      *time.Time `gorm:"default:now()"`
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
	var ids []int
	var clients []ClientTable
	dataFromFile, ids = readCSVFile(fileName)
	fmt.Println(dataFromFile, ids)
	insertNewRecords(dataFromFile, clients)
}

func readCSVFile(fileName multipart.File) ([][]string, []int) {

	//file, handler, err := fileName // получаем файл из формы

	defer fileName.Close()

	//path, err := os.Getwd()
	//file, err := os.Open(path + fileName)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

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
	return clientsFromFile, ids
}
func runExportClients(clients []ClientTable) string {
	date := time.Now().Format("2006-01-02 15:04:05")
	fileName := fmt.Sprintf("%x.csv", md5.Sum([]byte(date)))

	file, err := os.Create("views/files/" + fileName)
	if err != nil {
		// Обработка ошибки создания файла
		return ""
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	// Записываем данные пользователей
	for _, user := range clients {

		var userVerify string
		if user.Verify != nil {
			userVerify = strconv.FormatBool(*user.Verify)
		}
		var userVerifyDT string
		if user.VerifyDt != nil {
			userVerifyDT = user.VerifyDt.Format("2006-01-02 15:04:05")
		}
		var BDay string
		if user.BDay != nil {
			BDay = user.BDay.Format("2006-01-02")
		}
		var RegDate string
		if user.RegDate != nil {
			RegDate = user.RegDate.Format("2006-01-02")
		}
		var GroupCreate string
		if user.GroupCreate != nil {
			GroupCreate = user.GroupCreate.Format("2006-01-02 15:04:05")
		}
		err := writer.Write([]string{
			strconv.Itoa(user.ID),
			strconv.Itoa(user.ClubID),
			user.Login,
			user.Password,
			user.Phone,
			user.Email,
			user.Icon,
			strconv.FormatFloat(user.Amount, 'f', 2, 64),
			strconv.FormatFloat(user.Bonus, 'f', 2, 64),
			strconv.Itoa(user.TotalTime),
			user.FullName,
			strconv.FormatBool(*user.StatusActive),
			user.TelegramID,
			user.VKID,
			RegDate,
			BDay,
			userVerify,
			userVerifyDT,
			user.MiddleName,
			user.Surname,
			user.Name,
			strconv.Itoa(user.GroupID),
			strconv.FormatFloat(user.GroupAmount, 'f', -1, 64),
			GroupCreate,
		})

		if err != nil {
			return ""
		}
	}
	return fileName
}

func insertNewRecords(clientsFromFile [][]string, clients []ClientTable) {
	for _, clientFromFile := range clientsFromFile {
		var client ClientTable
		err := Database.Where("login = ?", clientFromFile[2]).First(&client).Error
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				// Создание нового клиента на основе данных из clientFromFile
				var groupCreate *time.Time
				if clientFromFile[23] != "" {
					groupCreateTime, _ := time.Parse("2006-01-02", clientFromFile[24])
					groupCreate = &groupCreateTime
				}
				clientFromFile1, _ := strconv.Atoi(clientFromFile[1])
				clientFromFile9, _ := strconv.Atoi(clientFromFile[9])
				clientFromFile7, _ := strconv.ParseFloat(clientFromFile[7], 64)
				clientFromFile8, _ := strconv.ParseFloat(clientFromFile[8], 64)
				statusActive := clientFromFile[11]
				temp, _ := strconv.ParseBool(statusActive)
				statusActiveBool := &temp
				var RegDate *time.Time
				if clientFromFile[14] != "" {
					RegDate = parseDateTime(clientFromFile[14])
				} else {
					RegDate = nil
				}
				var Verify *bool
				if clientFromFile[16] != "" {
					temp2, _ := strconv.ParseBool(clientFromFile[16])
					Verify = &temp2
				}
				GroupId, _ := strconv.Atoi(clientFromFile[21])
				GroupAmount, _ := strconv.ParseFloat(clientFromFile[22], 64)

				client := ClientTable{
					ClubID:       clientFromFile1,
					Login:        clientFromFile[2],
					Password:     clientFromFile[3],
					Phone:        clientFromFile[4],
					Email:        clientFromFile[5],
					Icon:         clientFromFile[6],
					Amount:       clientFromFile7,
					Bonus:        clientFromFile8,
					TotalTime:    clientFromFile9,
					FullName:     clientFromFile[10],
					StatusActive: statusActiveBool,
					TelegramID:   clientFromFile[12],
					VKID:         clientFromFile[13],
					RegDate:      RegDate,
					BDay:         parseDateTime(clientFromFile[15]),
					Verify:       Verify,
					VerifyDt:     parseDateTime(clientFromFile[17]),
					Name:         clientFromFile[18],
					Surname:      clientFromFile[19],
					MiddleName:   clientFromFile[20],
					GroupID:      GroupId,
					GroupAmount:  GroupAmount,
					GroupCreate:  groupCreate,
				}
				Database.Create(&client)
			}
		}
	}
}

func parseDateTime(dateTimeStr string) *time.Time {
	if dateTimeStr == "" {
		return nil
	}
	parsedDateTime, _ := time.Parse("2006-01-02 15:04:05", dateTimeStr)
	return &parsedDateTime
}
