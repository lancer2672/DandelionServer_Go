package service

import (
	"context"
	"database/sql"

	db "github.com/lancer2672/DandelionServer_Go/db/sqlc"
)

type MovieService struct {
	store *db.Store
}

type MovieServiceInterface interface {
	GetMovie(ctx context.Context, id int32) (db.Movie, error)
	GetMoviesByGenre(ctx context.Context, genreID int32, limit int64, offset int64) ([]db.GetMoviesByGenreRow, error)
	GetMoviesBySeries(ctx context.Context, seriesID sql.NullInt32, limit int64, offset int64) ([]db.Movie, error)
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
func (s *MovieService) GetMoviesByGenre(ctx context.Context, genreID int32, limit int64, offset int64) ([]db.GetMoviesByGenreRow, error) {
	movies, err := s.store.GetMoviesByGenre(ctx, db.GetMoviesByGenreParams{
		GenreID: genreID,
		Limit:   limit,
		Offset:  offset,
	})
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *MovieService) GetMoviesBySeries(ctx context.Context, seriesID sql.NullInt32, limit int64, offset int64) ([]db.Movie, error) {
	movies, err := s.store.GetMoviesBySeries(ctx, db.GetMoviesBySeriesParams{
		SeriesID: seriesID,
		Limit:    limit,
		Offset:   offset,
	})
	if err != nil {
		return nil, err
	}
	return movies, nil
}
