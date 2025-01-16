package models

import "time"

type Order struct {
	ID      int       `gorm:"primary_key"`
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
	Status  string    `json:"status"`
	Latlng  string    `json:"latlng"`
	Date    time.Time `json:"date"`
	StoreId int       `json:"storeId"`
	Uid     int       `json:"uid"`
	Items   string    `json:"items"`
}

func (u *Order) TableName() string {
	return "orders"
}
