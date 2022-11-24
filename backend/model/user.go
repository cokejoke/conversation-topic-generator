package model

type User struct {
	ID       int    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name     string `json:"name" sql:"index"`
	Username string `json:"username"`
	Password string `json:"password"`
}
