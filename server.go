package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"refresh/models"
)

type Config struct {
	Db struct {
		ConnectionString string `json:"connectionString"`
	} `json:"database"`
}

func main() {

	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	dsn := config.Db.ConnectionString
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	models.Database = db
	if err != nil {
		panic("failed to connect database")
	}

	StartRouter()

	fmt.Println("Server is listening http://localhost:8185")
	http.ListenAndServe(":8185", nil)
}
