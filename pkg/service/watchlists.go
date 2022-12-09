package service

import (
	"crypto_app/pkg/models"
	"crypto_app/pkg/repository"
)

type WatchlistsService struct {
	repository repository.Watchlists
}

func NewWatchlistsService(repository repository.Watchlists) *WatchlistsService {
	return &WatchlistsService{repository: repository}
}

func (s *WatchlistsService) CreateWatchlist(watchlist models.Watchlist) error {
	return s.repository.CreateWatchlist(watchlist)
}
