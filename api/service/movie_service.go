package service

import (
	"context"

	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
)

type MovieService struct {
	store *db.Store
}

type MovieServiceInterface interface {
	GetMovie(ctx context.Context, id int32) (db.Movie, error)
}

func NewMovieService(store *db.Store) *MovieService {
	return &MovieService{
		store: store,
	}
}

func (s *MovieService) GetMovie(ctx context.Context, id int32) (db.Movie, error) {
	movie, err := s.store.GetMovie(ctx, id)
	if err != nil {
		return db.Movie{}, err
	}
	return movie, nil
}
