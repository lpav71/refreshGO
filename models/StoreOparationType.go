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

func (StoreOperationType) GetStoreOperationType(w http.ResponseWriter, r *http.Request) {
	var storeOperationTypes []StoreOperationType

	// Запрос на получение типов операций
	if err := Database.Find(&storeOperationTypes).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(storeOperationTypes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
