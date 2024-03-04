package service

import (
	"context"
	"fmt"

	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
)

type MovieService struct {
	store *db.Store
}

type MovieServiceInterface interface {
	GetMovie(ctx context.Context, id int32) (db.Movie, error)
	GetMoviesByGenre(ctx context.Context, arg db.GetMoviesByGenreParams) ([]db.GetMoviesByGenreRow, error)
	GetMoviesBySerie(ctx context.Context, arg db.GetMoviesBySeriesParams) ([]db.GetMoviesBySeriesRow, error)
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

func (s *MovieService) GetMoviesByGenre(ctx context.Context, arg db.GetMoviesByGenreParams) ([]db.GetMoviesByGenreRow, error) {
	fmt.Println(arg)
	movies, err := s.store.GetMoviesByGenre(ctx, arg)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *MovieService) GetMoviesBySerie(ctx context.Context, arg db.GetMoviesBySeriesParams) ([]db.GetMoviesBySeriesRow, error) {
	movies, err := s.store.GetMoviesBySerie(ctx, arg)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
