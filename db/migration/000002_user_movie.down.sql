ALTER TABLE "user_movies" DROP CONSTRAINT IF EXISTS "user_movies_user_id_fkey";
ALTER TABLE "user_movies" DROP CONSTRAINT IF EXISTS "user_movies_movie_id_fkey";

DROP TABLE IF EXISTS "user_movies";