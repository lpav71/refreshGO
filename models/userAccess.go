package models

import (
	"encoding/json"
	"net/http"
)

type UserAccess struct {
	ClubID         int
	UserID         int64
	NewUser        bool
	BalanceMoney   bool
	BalanceBonus   bool
	RetMoney       bool
	RetBonus       bool
	OpenShift      bool
	CloseShift     bool
	CreatePC       bool
	CreateZone     bool
	CreateTask     bool
	ShopPay        bool
	CreatePromo    bool
	CreateGame     bool
	DeleteGame     bool `gorm:"column:del_game"`
	EditGame       bool
	CreateSteamAcc bool
	PriceCreate    bool
	PriceDelete    bool `gorm:"column:price_del"`
}

func (UserAccess) TableName() string {
	return "user_access"
}

func (UserAccess) GetPermisions(w http.ResponseWriter, r *http.Request) {
	var userAccess UserAccess
	r.ParseForm()
	clubId := r.Form.Get("club_id")
	userId := r.Form.Get("user_id")
	Database.First(&userAccess, "club_id = ? AND user_id = ?", clubId, userId)
	jsonData, _ := json.Marshal(userAccess)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
