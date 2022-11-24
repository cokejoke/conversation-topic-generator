package model

type Pack struct {
	ID       int     `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name     string  `json:"name" validate:"required,min=1"`
	Author   string  `json:"author" validate:"required,min=1"`
	ImageUrl string  `json:"imageUrl"`
	Language string  `json:"language" validate:"required,min=1"`
	Topics   []Topic `json:"topics"` // one-to-many relationship

}
