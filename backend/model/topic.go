package model

type Topic struct {
	ID     int    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Topic  string `json:"topic"`
	PackID int    `json:"packID"`
}
