package models

type Watchlist struct {
	Id             int    `json:"id" db:"watchlist_id"`
	UserId         int    `json:"user_id" db:"user_id"`
	Watchlist      string `json:"watchlist" db:"watchlist" binding:"required"`
	CryptoCurrency string `json:"crypto_currency" db:"crypto_currency" binding:"required"`
}
