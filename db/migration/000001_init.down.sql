
ALTER TABLE "votes" DROP CONSTRAINT IF EXISTS "votes_movie_id_fkey";
ALTER TABLE "votes" DROP CONSTRAINT IF EXISTS "votes_user_id_fkey";

ALTER TABLE "rooms" DROP CONSTRAINT IF EXISTS "rooms_movie_id_fkey";
ALTER TABLE "movie_genres" DROP CONSTRAINT IF EXISTS "movie_genres_genre_id_fkey";
ALTER TABLE "movie_genres" DROP CONSTRAINT IF EXISTS "movie_genres_movie_id_fkey";
ALTER TABLE "movie_history" DROP CONSTRAINT IF EXISTS "movie_history_users_id_fkey";
ALTER TABLE "movie_history" DROP CONSTRAINT IF EXISTS "movie_history_movies_id_fkey";

-- Drop tables
DROP TABLE IF EXISTS "movie_history";
DROP TABLE IF EXISTS "votes";
DROP TABLE IF EXISTS "rooms";
DROP TABLE IF EXISTS "movie_genres";
DROP TABLE IF EXISTS "movies";
DROP TABLE IF EXISTS "series";
DROP TABLE IF EXISTS "genres";
DROP TABLE IF EXISTS "users";
