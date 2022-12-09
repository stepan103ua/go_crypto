package repository

import (
	"crypto_app/pkg/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type WatchlistsPostgres struct {
	db *sqlx.DB
}

func NewWatchlistsPostgres(db *sqlx.DB) *WatchlistsPostgres {
	return &WatchlistsPostgres{db}
}

func (r *WatchlistsPostgres) CreateWatchlist(watchlist models.Watchlist) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, watchlist, crypto_currency) values ($1, $2, $3)", watchlistTable)

	_, err := r.db.Exec(query, watchlist.UserId, watchlist.Watchlist, watchlist.CryptoCurrency)

	return err
}

func (r *WatchlistsPostgres) GetAllUserWatchlists(userId int) ([]models.Watchlist, error) {
	var allWatchlists []models.Watchlist

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", watchlistTable)

	err := r.db.Select(&allWatchlists, query, userId)

	return allWatchlists, err
}
