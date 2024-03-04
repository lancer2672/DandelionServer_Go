// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.getListGenresStmt, err = db.PrepareContext(ctx, getListGenres); err != nil {
		return nil, fmt.Errorf("error preparing query GetListGenres: %w", err)
	}
	if q.getListMoviesStmt, err = db.PrepareContext(ctx, getListMovies); err != nil {
		return nil, fmt.Errorf("error preparing query GetListMovies: %w", err)
	}
	if q.getMovieStmt, err = db.PrepareContext(ctx, getMovie); err != nil {
		return nil, fmt.Errorf("error preparing query GetMovie: %w", err)
	}
	if q.getMoviesByGenreStmt, err = db.PrepareContext(ctx, getMoviesByGenre); err != nil {
		return nil, fmt.Errorf("error preparing query GetMoviesByGenre: %w", err)
	}
	if q.getMoviesBySeriesStmt, err = db.PrepareContext(ctx, getMoviesBySeries); err != nil {
		return nil, fmt.Errorf("error preparing query GetMoviesBySerie: %w", err)
	}
	if q.getRecentMoviesStmt, err = db.PrepareContext(ctx, getRecentMovies); err != nil {
		return nil, fmt.Errorf("error preparing query GetRecentMovies: %w", err)
	}
	if q.getRoomStmt, err = db.PrepareContext(ctx, getRoom); err != nil {
		return nil, fmt.Errorf("error preparing query GetRoom: %w", err)
	}
	if q.getRoomsByUserStmt, err = db.PrepareContext(ctx, getRoomsByUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetRoomsByUser: %w", err)
	}
	if q.getVotesByUserStmt, err = db.PrepareContext(ctx, getVotesByUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetVotesByUser: %w", err)
	}
	if q.searchMoviesStmt, err = db.PrepareContext(ctx, searchMovies); err != nil {
		return nil, fmt.Errorf("error preparing query SearchMovies: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.getListGenresStmt != nil {
		if cerr := q.getListGenresStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListGenresStmt: %w", cerr)
		}
	}
	if q.getListMoviesStmt != nil {
		if cerr := q.getListMoviesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListMoviesStmt: %w", cerr)
		}
	}
	if q.getMovieStmt != nil {
		if cerr := q.getMovieStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMovieStmt: %w", cerr)
		}
	}
	if q.getMoviesByGenreStmt != nil {
		if cerr := q.getMoviesByGenreStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMoviesByGenreStmt: %w", cerr)
		}
	}
	if q.getMoviesBySeriesStmt != nil {
		if cerr := q.getMoviesBySeriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMoviesBySeriesStmt: %w", cerr)
		}
	}
	if q.getRecentMoviesStmt != nil {
		if cerr := q.getRecentMoviesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRecentMoviesStmt: %w", cerr)
		}
	}
	if q.getRoomStmt != nil {
		if cerr := q.getRoomStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRoomStmt: %w", cerr)
		}
	}
	if q.getRoomsByUserStmt != nil {
		if cerr := q.getRoomsByUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRoomsByUserStmt: %w", cerr)
		}
	}
	if q.getVotesByUserStmt != nil {
		if cerr := q.getVotesByUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getVotesByUserStmt: %w", cerr)
		}
	}
	if q.searchMoviesStmt != nil {
		if cerr := q.searchMoviesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing searchMoviesStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                    DBTX
	tx                    *sql.Tx
	getListGenresStmt     *sql.Stmt
	getListMoviesStmt     *sql.Stmt
	getMovieStmt          *sql.Stmt
	getMoviesByGenreStmt  *sql.Stmt
	getMoviesBySeriesStmt *sql.Stmt
	getRecentMoviesStmt   *sql.Stmt
	getRoomStmt           *sql.Stmt
	getRoomsByUserStmt    *sql.Stmt
	getVotesByUserStmt    *sql.Stmt
	searchMoviesStmt      *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                    tx,
		tx:                    tx,
		getListGenresStmt:     q.getListGenresStmt,
		getListMoviesStmt:     q.getListMoviesStmt,
		getMovieStmt:          q.getMovieStmt,
		getMoviesByGenreStmt:  q.getMoviesByGenreStmt,
		getMoviesBySeriesStmt: q.getMoviesBySeriesStmt,
		getRecentMoviesStmt:   q.getRecentMoviesStmt,
		getRoomStmt:           q.getRoomStmt,
		getRoomsByUserStmt:    q.getRoomsByUserStmt,
		getVotesByUserStmt:    q.getVotesByUserStmt,
		searchMoviesStmt:      q.searchMoviesStmt,
	}
}
