package models

import (
	"encoding/json"
	"net/http"
)

type StoreOperationType struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar"`
}

func (StoreOperationType) TableName() string {
	return "store_operation_type"
}

func (StoreOperationType) GetStoreOparationType(w http.ResponseWriter, r *http.Request) {
	var storeOperationTypes []StoreOperationType
	Database.Find(&storeOperationTypes)
	jsonData, _ := json.Marshal(storeOperationTypes)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
