
ALTER TABLE votes ALTER COLUMN stars DROP NOT NULL;
ALTER TABLE votes ALTER COLUMN movie_id DROP NOT NULL;
ALTER TABLE votes ALTER COLUMN user_id DROP NOT NULL;
ALTER TABLE votes ALTER COLUMN id DROP NOT NULL;

ALTER TABLE rooms ALTER COLUMN created_at DROP NOT NULL;
ALTER TABLE rooms ALTER COLUMN created_by DROP NOT NULL;
ALTER TABLE rooms ALTER COLUMN movie_id DROP NOT NULL;
ALTER TABLE rooms ALTER COLUMN id DROP NOT NULL;

ALTER TABLE movie_genres ALTER COLUMN genre_id DROP NOT NULL;
ALTER TABLE movie_genres ALTER COLUMN movie_id DROP NOT NULL;

ALTER TABLE movies ALTER COLUMN created_at DROP NOT NULL;
ALTER TABLE movies ALTER COLUMN stars DROP NOT NULL;
ALTER TABLE movies ALTER COLUMN views DROP NOT NULL;
ALTER TABLE movies ALTER COLUMN thumbnail DROP NOT NULL;
ALTER TABLE movies ALTER COLUMN file_path DROP NOT NULL;
ALTER TABLE movies ALTER COLUMN trailer DROP NOT NULL;
ALTER TABLE movies ALTER COLUMN actor_avatars DROP NOT NULL;
ALTER TABLE movies ALTER COLUMN description DROP NOT NULL;
ALTER TABLE movies ALTER COLUMN duration DROP NOT NULL;
ALTER TABLE movies ALTER COLUMN title DROP NOT NULL;
ALTER TABLE movies ALTER COLUMN id DROP NOT NULL;

ALTER TABLE series ALTER COLUMN thumbnail DROP NOT NULL;
ALTER TABLE series ALTER COLUMN name DROP NOT NULL;
ALTER TABLE series ALTER COLUMN id DROP NOT NULL;

ALTER TABLE genres ALTER COLUMN name DROP NOT NULL;
ALTER TABLE genres ALTER COLUMN id DROP NOT NULL;

ALTER TABLE users ALTER COLUMN username DROP NOT NULL;
ALTER TABLE users ALTER COLUMN id DROP NOT NULL;

-- Add the series_id column back to the movies table
ALTER TABLE movies ADD COLUMN series_id INTEGER;

-- Drop the movies_series table
DROP TABLE IF EXISTS "movies_series";
