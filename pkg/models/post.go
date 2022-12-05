package models

import "time"

type Post struct {
	Id             int       `json:"id" db:"id"`
	Title          string    `json:"title" binding:"required" db:"title"`
	Description    string    `json:"description" binding:"required" db:"description"`
	ImageUrl       string    `json:"image_url" binding:"required" db:"image_url"`
	CryptoCurrency string    `json:"crypto_currency" binding:"required" db:"crypto_currency"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UserId         int       `json:"owner_id" db:"owner_id"`
}
