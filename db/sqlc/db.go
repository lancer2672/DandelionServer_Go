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
	if q.createMovieStmt, err = db.PrepareContext(ctx, createMovie); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMovie: %w", err)
	}
	if q.createMovieHistoryStmt, err = db.PrepareContext(ctx, createMovieHistory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMovieHistory: %w", err)
	}
	if q.getListGenresStmt, err = db.PrepareContext(ctx, getListGenres); err != nil {
		return nil, fmt.Errorf("error preparing query GetListGenres: %w", err)
	}
	if q.getListMoviesStmt, err = db.PrepareContext(ctx, getListMovies); err != nil {
		return nil, fmt.Errorf("error preparing query GetListMovies: %w", err)
	}
	if q.getMovieStmt, err = db.PrepareContext(ctx, getMovie); err != nil {
		return nil, fmt.Errorf("error preparing query GetMovie: %w", err)
	}
	if q.getMovieHistoryByUserIdStmt, err = db.PrepareContext(ctx, getMovieHistoryByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query GetMovieHistoryByUserId: %w", err)
	}
	if q.getMoviesByGenreStmt, err = db.PrepareContext(ctx, getMoviesByGenre); err != nil {
		return nil, fmt.Errorf("error preparing query GetMoviesByGenre: %w", err)
	}
	if q.getMoviesBySerieStmt, err = db.PrepareContext(ctx, getMoviesBySerie); err != nil {
		return nil, fmt.Errorf("error preparing query GetMoviesBySerie: %w", err)
	}
	if q.getRecentMoviesStmt, err = db.PrepareContext(ctx, getRecentMovies); err != nil {
		return nil, fmt.Errorf("error preparing query GetRecentMovies: %w", err)
	}
	if q.getVotesByUserStmt, err = db.PrepareContext(ctx, getVotesByUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetVotesByUser: %w", err)
	}
	if q.searchMoviesStmt, err = db.PrepareContext(ctx, searchMovies); err != nil {
		return nil, fmt.Errorf("error preparing query SearchMovies: %w", err)
	}
	if q.updateMovieHistoryStmt, err = db.PrepareContext(ctx, updateMovieHistory); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMovieHistory: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createMovieStmt != nil {
		if cerr := q.createMovieStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMovieStmt: %w", cerr)
		}
	}
	if q.createMovieHistoryStmt != nil {
		if cerr := q.createMovieHistoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMovieHistoryStmt: %w", cerr)
		}
	}
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
	if q.getMovieHistoryByUserIdStmt != nil {
		if cerr := q.getMovieHistoryByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMovieHistoryByUserIdStmt: %w", cerr)
		}
	}
	if q.getMoviesByGenreStmt != nil {
		if cerr := q.getMoviesByGenreStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMoviesByGenreStmt: %w", cerr)
		}
	}
	if q.getMoviesBySerieStmt != nil {
		if cerr := q.getMoviesBySerieStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMoviesBySerieStmt: %w", cerr)
		}
	}
	if q.getRecentMoviesStmt != nil {
		if cerr := q.getRecentMoviesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getRecentMoviesStmt: %w", cerr)
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
	if q.updateMovieHistoryStmt != nil {
		if cerr := q.updateMovieHistoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMovieHistoryStmt: %w", cerr)
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
	db                          DBTX
	tx                          *sql.Tx
	createMovieStmt             *sql.Stmt
	createMovieHistoryStmt      *sql.Stmt
	getListGenresStmt           *sql.Stmt
	getListMoviesStmt           *sql.Stmt
	getMovieStmt                *sql.Stmt
	getMovieHistoryByUserIdStmt *sql.Stmt
	getMoviesByGenreStmt        *sql.Stmt
	getMoviesBySerieStmt        *sql.Stmt
	getRecentMoviesStmt         *sql.Stmt
	getVotesByUserStmt          *sql.Stmt
	searchMoviesStmt            *sql.Stmt
	updateMovieHistoryStmt      *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                          tx,
		tx:                          tx,
		createMovieStmt:             q.createMovieStmt,
		createMovieHistoryStmt:      q.createMovieHistoryStmt,
		getListGenresStmt:           q.getListGenresStmt,
		getListMoviesStmt:           q.getListMoviesStmt,
		getMovieStmt:                q.getMovieStmt,
		getMovieHistoryByUserIdStmt: q.getMovieHistoryByUserIdStmt,
		getMoviesByGenreStmt:        q.getMoviesByGenreStmt,
		getMoviesBySerieStmt:        q.getMoviesBySerieStmt,
		getRecentMoviesStmt:         q.getRecentMoviesStmt,
		getVotesByUserStmt:          q.getVotesByUserStmt,
		searchMoviesStmt:            q.searchMoviesStmt,
		updateMovieHistoryStmt:      q.updateMovieHistoryStmt,
	}
}
