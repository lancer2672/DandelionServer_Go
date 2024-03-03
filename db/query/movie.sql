-- name: GetMoviesByGenre :many
SELECT * FROM movies
JOIN movie_genres ON movies.id = movie_genres.movie_id
WHERE movie_genres.genre_id = $1
ORDER BY movies.id
LIMIT $2
OFFSET $3;

-- name: GetMoviesBySeries :many
SELECT * FROM movies
WHERE series_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: GetMovie :one
SELECT * FROM movies
WHERE id = $1 LIMIT 1;

-- name: ListMovies :many
SELECT * FROM movies
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetMoviesByUser :many
SELECT * FROM movies
JOIN user_rooms ON movies.id = user_movies.movie_id
WHERE user_rooms.user_id = $1
ORDER BY movies.id;

-- name: GetRecentMovies :many
SELECT * FROM movies
ORDER BY movies.created_at DESC
LIMIT $1;

-- name: SearchMovies :many
SELECT * FROM movies
WHERE movies.title LIKE '%' || $1 || '%'
ORDER BY movies.id
LIMIT $2
OFFSET $3;

