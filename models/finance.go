package models

import (
	"encoding/json"
	"net/http"
	"time"
)

type Finance struct {
	ID              uint       `gorm:"primaryKey"`
	ClubID          int        `gorm:"not null"`
	AdminID         int        `gorm:"not null"`
	DtCreate        time.Time  `gorm:"not null;default:now()"`
	Shift           int        `gorm:"null"`
	OpenShift       *time.Time `gorm:"null"`
	CloseShift      *time.Time `gorm:"null"`
	Cash            float64    `gorm:"not null;default:0"`
	CashNum         int        `gorm:"not null;default:0"`
	NoCash          float64    `gorm:"not null;default:0"`
	NoCashNum       int        `gorm:"not null;default:0"`
	ReturnCash      float64    `gorm:"not null;default:0"`
	ReturnCashNum   int        `gorm:"not null;default:0"`
	ReturnNoCash    float64    `gorm:"not null;default:0"`
	ReturnNoCashNum int        `gorm:"not null;default:0"`
	Bonus           float64    `gorm:"not null;default:0"`
	BonusNum        int        `gorm:"not null;default:0"`
	Status          bool       `gorm:"not null;default:true"`
	ShopCash        float64    `gorm:"not null;default:0"`
	ShopCashNum     float64    `gorm:"not null;default:0"`
	ShopNoCash      float64    `gorm:"not null;default:0"`
	ShopNoCashNum   float64    `gorm:"not null;default:0"`
	CashBoxSerial   *string    `gorm:"null"`
	Time            *string
}

type Data struct {
	Shift   string
	Finance []Finance
}

func (Finance) TableName() string {
	return "finance"
}

func (Finance) GetSmena(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clubId := r.Form.Get("club_id")

	var count int64
	Database.Table("finance").
		Select("cash, nocash, shop_cash, shop_nocash, to_char(open_shift, 'DD.MM.YYYY HH24:MI') AS time").
		Where("status = true").
		Where("club_id = ?", clubId).
		Count(&count)

	shift := "open"

	var finance []Finance
	Database.Table("finance").
		Select("cash, nocash, shop_cash, shop_nocash, to_char(open_shift, 'DD.MM.YYYY HH24:MI') AS time").
		Where("status = true").
		Where("club_id = ?", clubId).
		Find(&finance)

	if count == 0 {
		Database.Where("status = ? AND club_id = ?", false, 1).
			Order("close_shift DESC").
			Select("cash, nocash, shop_cash, shop_nocash, to_char(close_shift, 'DD.MM.YYYY HH24:MI') AS time").
			Find(&finance)
		shift = "close"
	}
	outData := Data{
		Shift:   shift,
		Finance: finance,
	}
	jsonData, _ := json.Marshal(outData)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
