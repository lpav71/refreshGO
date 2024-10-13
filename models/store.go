package models

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Store struct {
	ID           int       `gorm:"column:id;primary_key"`
	ClubID       int       `gorm:"column:club_id"`
	AdminID      int       `gorm:"column:admin_id"`
	AdminName    string    `gorm:"column:admin_name"`
	DtCreate     time.Time `gorm:"column:dt_create"`
	Product      string    `gorm:"column:product"`
	ProductParam string    `gorm:"column:product_param"`
	Barcode      string    `gorm:"column:barcode"`
	ShellShow    bool      `gorm:"column:shell_show"`
	Icon         string    `gorm:"column:icon"`
	Num          int       `gorm:"column:num"`
	Types        int       `gorm:"column:types"`
	Price        float64   `gorm:"column:price"`
	PriceBonus   float64   `gorm:"column:price_bonus"`
	Discount     bool      `gorm:"column:discount"`
	Nalog        int       `gorm:"column:nalog"`
	Name         string
}

type ProductType struct {
	ID       int
	ClubID   int
	Status   bool
	Name     string
	Types    int
	Category int
}

func (Store) TableName() string {
	return "store"
}

func (ProductType) TableName() string {
	return "product_type"
}

func (Store) GetAll(w http.ResponseWriter, r *http.Request) {
	clubID, _ := strconv.Atoi(r.FormValue("club_id"))

	storeChan := make(chan []Store)
	productTypeChan := make(chan []ProductType)

	go getAllStore(clubID, storeChan)
	go getProductType(clubID, productTypeChan)

	store := <-storeChan
	productType := <-productTypeChan

	outData := []interface{}{store, productType}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(outData)
}

func getAllStore(clubID int, storeChan chan []Store) {
	var stores []Store

	err := Database.Table("store").
		Select("store.id as storeid, store.*").
		Joins("JOIN product_type ON store.types = product_type.types").
		Where("store.club_id = ?", clubID).
		Order("storeid").
		Find(&stores).Error

	if err != nil {
		storeChan <- nil
		return
	}

	storeChan <- stores
}

func getProductType(clubID int, productTypeChan chan []ProductType) {
	var productTypes []ProductType
	Database.Where("club_id = ?", clubID).Find(&productTypes)
	productTypeChan <- productTypes
}
