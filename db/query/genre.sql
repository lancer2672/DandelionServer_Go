-- name: ListGenres :many
SELECT * FROM genres
ORDER BY id
LIMIT $1
OFFSET $2;
