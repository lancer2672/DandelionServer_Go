-- name: CreateMovieHistory :exec
INSERT INTO movie_history
 (user_id,movie_id,watched_duration, last_watched)
VALUES ($1, $2, $3, $4);
 
-- name: UpdateMovieHistory :exec
UPDATE movie_history
SET watched_duration = $1, last_watched = $2
WHERE movie_id = $3;