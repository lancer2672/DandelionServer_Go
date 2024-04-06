
-- name: GetMovie :one
WITH movie_genres AS (
  SELECT movie_id, ARRAY_AGG(genres.name) AS genres
  FROM movie_genres
  JOIN genres ON movie_genres.genre_id = genres.id
  GROUP BY movie_id
)
SELECT movies.*, movie_genres.genres FROM movies
LEFT JOIN movie_genres ON movies.id = movie_genres.movie_id
WHERE movies.id = $1 LIMIT 1;

-- name: GetMoviesByGenre :many
WITH movie_genres AS (
  SELECT movie_id, ARRAY_AGG(genres.name) AS genres
  FROM movie_genres
  JOIN genres ON movie_genres.genre_id = genres.id
  GROUP BY movie_id
)
SELECT movies.*, movie_genres.genres FROM movies
JOIN movie_genres ON movies.id = movie_genres.movie_id
WHERE movie_genres.genre_id = $1
ORDER BY movies.id
LIMIT $2
OFFSET $3;

-- name: GetMoviesBySerie :many
WITH movie_genres AS (
  SELECT movie_id, ARRAY_AGG(genres.name) AS genres
  FROM movie_genres
  JOIN genres ON movie_genres.genre_id = genres.id
  GROUP BY movie_id
)
SELECT movies.*, movie_genres.genres FROM movies
JOIN movies_series ON movies.id = movies_series.movie_id
LEFT JOIN movie_genres ON movies.id = movie_genres.movie_id
WHERE movies_series.series_id = $1
ORDER BY movies.id
LIMIT $2
OFFSET $3;

-- name: GetListMovies :many
WITH movie_genres AS (
  SELECT movie_id, ARRAY_AGG(genres.name) AS genres
  FROM movie_genres
  JOIN genres ON movie_genres.genre_id = genres.id
  GROUP BY movie_id
)
SELECT movies.*, movie_genres.genres FROM movies
LEFT JOIN movie_genres ON movies.id = movie_genres.movie_id
ORDER BY movies.id
LIMIT $1
OFFSET $2;

-- name: GetRecentMovies :many
WITH movie_genres AS (
  SELECT movie_id, ARRAY_AGG(genres.name) AS genres
  FROM movie_genres
  JOIN genres ON movie_genres.genre_id = genres.id
  GROUP BY movie_id
)
SELECT movies.*, movie_genres.genres FROM movies
LEFT JOIN movie_genres ON movies.id = movie_genres.movie_id
ORDER BY movies.created_at DESC
LIMIT $1
OFFSET $2;

-- name: SearchMovies :many
WITH movie_genres AS (
  SELECT movie_id, ARRAY_AGG(genres.name) AS genres
  FROM movie_genres
  JOIN genres ON movie_genres.genre_id = genres.id
  GROUP BY movie_id
)
SELECT movies.*, movie_genres.genres FROM movies
LEFT JOIN movie_genres ON movies.id = movie_genres.movie_id
WHERE movies.title LIKE '%' || $1 || '%'
ORDER BY movies.id
LIMIT $2
OFFSET $3;

-- name: GetWatchingMovies :many
WITH movie_genres AS (
  SELECT movie_id, ARRAY_AGG(genres.name) AS genres
  FROM movie_genres
  JOIN genres ON movie_genres.genre_id = genres.id
  GROUP BY movie_id
)
SELECT movies.*, movie_genres.genres, movie_history.watched_duration, movie_history.last_watched FROM movies
JOIN movie_history ON movies.id = movie_history.movie_id
LEFT JOIN movie_genres ON movies.id = movie_genres.movie_id
WHERE movie_history.user_id = $1 AND (movie_history.watched_duration / movies.duration) > 0.1
ORDER BY movie_history.last_watched DESC
LIMIT $2
OFFSET $3;
