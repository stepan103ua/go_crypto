package models

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"name" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

type UserUpdate struct {
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
