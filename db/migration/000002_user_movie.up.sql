
CREATE TABLE "user_movies" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "movie_id" integer,
  "upload_time" timestamp DEFAULT 'now()'
);

ALTER TABLE "user_movies" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_movies" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");
