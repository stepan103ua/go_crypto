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
