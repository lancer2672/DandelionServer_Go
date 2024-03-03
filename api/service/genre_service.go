package service

import (
	"context"
	"fmt"

	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
)

type GenreService struct {
	store *db.Store
}

type GenreServiceInterface interface {
	GetListGenres(ctx context.Context, arg db.GetListGenresParams) ([]db.Genre, error)
}

func NewGenreService(store *db.Store) *GenreService {
	return &GenreService{
		store: store,
	}
}

func (s *GenreService) GetListGenres(ctx context.Context, arg db.GetListGenresParams) ([]db.Genre, error) {
	fmt.Println(arg)
	genres, err := s.store.GetListGenres(ctx, arg)
	if err != nil {
		return nil, err
	}
	return genres, nil
}
