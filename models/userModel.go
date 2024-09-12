package models

type User struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique"` 
	Password string `json:"password"`
}