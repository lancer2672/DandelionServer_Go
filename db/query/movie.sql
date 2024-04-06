
-- name: UpdateMovie :one
UPDATE movies 
SET 
    title = COALESCE(sqlc.narg('title'), title),
    description = COALESCE(sqlc.narg('description'), description),
    file_path = COALESCE(sqlc.narg(file_path), file_path),
    thumbnail = COALESCE(sqlc.narg(thumbnail), thumbnail)
WHERE 
    id = sqlc.arg(id)
RETURNING *;

-- name: CreateMovie :one
INSERT INTO movies (title, duration, description, actor_avatars, trailer, file_path, thumbnail, views, stars)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;
-- name: GetMovie :one
SELECT movies.*, 
       (SELECT json_agg(json_build_object('id', genres.id, 'name', genres.name)) FROM movie_genres JOIN genres ON movie_genres.genre_id = genres.id WHERE movies.id = movie_genres.movie_id) AS genres,
       (SELECT json_agg(json_build_object('id', series.id, 'name', series.name)) FROM movies_series JOIN series ON movies_series.series_id = series.id WHERE movies.id = movies_series.movie_id) AS series
FROM movies
WHERE movies.id = $1
LIMIT 1;

-- name: GetMoviesByGenre :many
SELECT movies.*, 
       (SELECT json_agg(json_build_object('id', genres.id, 'name', genres.name)) FROM movie_genres JOIN genres ON movie_genres.genre_id = genres.id WHERE movies.id = movie_genres.movie_id) AS genres,
       (SELECT json_agg(json_build_object('id', series.id, 'name', series.name)) FROM movies_series JOIN series ON movies_series.series_id = series.id WHERE movies.id = movies_series.movie_id) AS series
FROM movies
JOIN movie_genres ON movies.id = movie_genres.movie_id
WHERE movie_genres.genre_id = $1
ORDER BY movies.id
LIMIT $2
OFFSET $3;

-- name: GetMoviesBySerie :many
SELECT movies.*, 
       (SELECT json_agg(json_build_object('id', genres.id, 'name', genres.name)) FROM movie_genres JOIN genres ON movie_genres.genre_id = genres.id WHERE movies.id = movie_genres.movie_id) AS genres,
       (SELECT json_agg(json_build_object('id', series.id, 'name', series.name)) FROM movies_series JOIN series ON movies_series.series_id = series.id WHERE movies.id = movies_series.movie_id) AS series
FROM movies
JOIN movies_series ON movies.id = movies_series.movie_id
WHERE movies_series.series_id = $1
ORDER BY movies.id
LIMIT $2
OFFSET $3;

-- name: GetListMovies :many
SELECT movies.*, 
       (SELECT json_agg(json_build_object('id', genres.id, 'name', genres.name)) FROM movie_genres JOIN genres ON movie_genres.genre_id = genres.id WHERE movies.id = movie_genres.movie_id) AS genres,
       (SELECT json_agg(json_build_object('id', series.id, 'name', series.name)) FROM movies_series JOIN series ON movies_series.series_id = series.id WHERE movies.id = movies_series.movie_id) AS series
FROM movies
ORDER BY movies.id
LIMIT $1
OFFSET $2;

-- name: GetRecentMovies :many
SELECT movies.*, 
       (SELECT json_agg(json_build_object('id', genres.id, 'name', genres.name)) FROM movie_genres JOIN genres ON movie_genres.genre_id = genres.id WHERE movies.id = movie_genres.movie_id) AS genres,
       (SELECT json_agg(json_build_object('id', series.id, 'name', series.name)) FROM movies_series JOIN series ON movies_series.series_id = series.id WHERE movies.id = movies_series.movie_id) AS series
FROM movies
ORDER BY movies.created_at DESC
LIMIT $1
OFFSET $2;

-- name: SearchMovies :many
SELECT movies.*, 
       (SELECT json_agg(json_build_object('id', genres.id, 'name', genres.name)) FROM movie_genres JOIN genres ON movie_genres.genre_id = genres.id WHERE movies.id = movie_genres.movie_id) AS genres,
       (SELECT json_agg(json_build_object('id', series.id, 'name', series.name)) FROM movies_series JOIN series ON movies_series.series_id = series.id WHERE movies.id = movies_series.movie_id) AS series
FROM movies
WHERE movies.title LIKE '%' || $1 || '%'
ORDER BY movies.id
LIMIT $2
OFFSET $3;
-- name: GetWatchingMovies :many
SELECT movies.*, 
       (SELECT json_agg(json_build_object('id', genres.id, 'name', genres.name)) FROM movie_genres JOIN genres ON movie_genres.genre_id = genres.id WHERE movies.id = movie_genres.movie_id) AS genres,
       (SELECT json_agg(json_build_object('id', series.id, 'name', series.name)) FROM movies_series JOIN series ON movies_series.series_id = series.id WHERE movies.id = movies_series.movie_id) AS series,
       movie_history.watched_duration, movie_history.last_watched
FROM movies
JOIN movie_history ON movies.id = movie_history.movie_id
WHERE movie_history.user_id = $1 AND (movie_history.watched_duration / movies.duration) > 0.1
ORDER BY movie_history.last_watched DESC
LIMIT $2
OFFSET $3;
