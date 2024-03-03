package service

import (
	"context"

	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
)

type GenreService struct {
	store *db.Store
}

type GenreServiceInterface interface {
	ListGenres(ctx context.Context, limit int64, offset int64) ([]db.Genre, error)
}

func NewGenreService(store *db.Store) *GenreService {
	return &GenreService{
		store: store,
	}
}

func (s *GenreService) ListGenres(ctx context.Context, limit int64, offset int64) ([]db.Genre, error) {
	genres, err := s.store.ListGenres(ctx, db.ListGenresParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	return genres, nil
}
