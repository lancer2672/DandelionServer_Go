// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: movie_history.sql

package database

import (
	"context"
	"time"
)

const createMovieHistory = `-- name: CreateMovieHistory :exec
INSERT INTO movie_history
 (user_id,movie_id,watched_duration, last_watched)
VALUES ($1, $2, $3, $4)
`

type CreateMovieHistoryParams struct {
	UserID          int32     `json:"user_id"`
	MovieID         int32     `json:"movie_id"`
	WatchedDuration int32     `json:"watched_duration"`
	LastWatched     time.Time `json:"last_watched"`
}

func (q *Queries) CreateMovieHistory(ctx context.Context, arg CreateMovieHistoryParams) error {
	_, err := q.exec(ctx, q.createMovieHistoryStmt, createMovieHistory,
		arg.UserID,
		arg.MovieID,
		arg.WatchedDuration,
		arg.LastWatched,
	)
	return err
}

const getMovieHistoryByUserId = `-- name: GetMovieHistoryByUserId :many
SELECT id, user_id, movie_id, watched_duration, last_watched FROM movie_history
WHERE user_id = $1
ORDER BY last_watched DESC
`

func (q *Queries) GetMovieHistoryByUserId(ctx context.Context, userID int32) ([]MovieHistory, error) {
	rows, err := q.query(ctx, q.getMovieHistoryByUserIdStmt, getMovieHistoryByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MovieHistory{}
	for rows.Next() {
		var i MovieHistory
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.MovieID,
			&i.WatchedDuration,
			&i.LastWatched,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMovieHistory = `-- name: UpdateMovieHistory :exec
UPDATE movie_history
SET watched_duration = $1, last_watched = $2
WHERE movie_id = $3
`

type UpdateMovieHistoryParams struct {
	WatchedDuration int32     `json:"watched_duration"`
	LastWatched     time.Time `json:"last_watched"`
	MovieID         int32     `json:"movie_id"`
}

func (q *Queries) UpdateMovieHistory(ctx context.Context, arg UpdateMovieHistoryParams) error {
	_, err := q.exec(ctx, q.updateMovieHistoryStmt, updateMovieHistory, arg.WatchedDuration, arg.LastWatched, arg.MovieID)
	return err
}
