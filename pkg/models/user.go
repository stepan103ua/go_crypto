package models

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"name" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
