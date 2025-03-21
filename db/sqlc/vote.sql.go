// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: vote.sql

package database

import (
	"context"
)

const getVotesByUser = `-- name: GetVotesByUser :many
SELECT id, user_id, movie_id, stars FROM votes
WHERE votes.user_id = $1
ORDER BY votes.id
`

func (q *Queries) GetVotesByUser(ctx context.Context, userID int32) ([]Vote, error) {
	rows, err := q.query(ctx, q.getVotesByUserStmt, getVotesByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vote{}
	for rows.Next() {
		var i Vote
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.MovieID,
			&i.Stars,
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
