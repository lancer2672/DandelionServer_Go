// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: room.sql

package database

import (
	"context"
	"time"
)

const getRoom = `-- name: GetRoom :one
SELECT id, movie_id, created_by, created_at FROM rooms
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRoom(ctx context.Context, id int32) (Room, error) {
	row := q.queryRow(ctx, q.getRoomStmt, getRoom, id)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.MovieID,
		&i.CreatedBy,
		&i.CreatedAt,
	)
	return i, err
}

const getRoomsByUser = `-- name: GetRoomsByUser :many
SELECT rooms.id, movie_id, created_by, created_at, user_rooms.id, user_id, room_id, joined_at FROM rooms
JOIN user_rooms ON rooms.id = user_rooms.room_id
WHERE user_rooms.user_id = $1
ORDER BY rooms.id
`

type GetRoomsByUserRow struct {
	ID        int32     `json:"id"`
	MovieID   int32     `json:"movie_id"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	ID_2      int32     `json:"id_2"`
	UserID    int32     `json:"user_id"`
	RoomID    int32     `json:"room_id"`
	JoinedAt  time.Time `json:"joined_at"`
}

func (q *Queries) GetRoomsByUser(ctx context.Context, userID int32) ([]GetRoomsByUserRow, error) {
	rows, err := q.query(ctx, q.getRoomsByUserStmt, getRoomsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetRoomsByUserRow{}
	for rows.Next() {
		var i GetRoomsByUserRow
		if err := rows.Scan(
			&i.ID,
			&i.MovieID,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.ID_2,
			&i.UserID,
			&i.RoomID,
			&i.JoinedAt,
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