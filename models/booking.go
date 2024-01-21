package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Booking struct {
	ID        int
	Phone     string
	Email     string
	Tariff    string
	Zone      string
	MapCompID int
	Login     string
	Date      string
	TimeStart string
	TimeStop  string
}

type Client struct {
	ID    int
	Phone string
	Email string
	Login string
}

type Price struct {
	ID   int
	Name string
}

func (Booking) GetAllClients(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clubId := r.Form.Get("club_id")
	var bookings []Booking
	err := Database.Table("booking").
		Select("booking.id, clients.phone, clients.email, price.name as tariff, zone.name as zone, booking.map_comp_id, clients.login, "+
			"to_char(booking.time_start, 'HH24:MI') as time_start, to_char(booking.time_stop, 'DD.MM.YYYY') as date, "+
			"to_char(booking.time_stop, 'HH24:MI') as time_stop").
		Joins("JOIN clients ON booking.user_id = clients.id").
		Joins("JOIN zone ON booking.id_zone = zone.id").
		Joins("JOIN price ON booking.price_id = price.id").
		Where("booking.club_id = ?", clubId).
		Where("clients.club_id = ?", clubId).
		Where("zone.club_id = ?", clubId).
		Where("price.club_id = ?", clubId).
		Where("booking.status = 0").
		Find(&bookings).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonData, err := json.Marshal(bookings)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
