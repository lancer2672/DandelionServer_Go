CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar
);

CREATE TABLE "genres" (
  "id" integer PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "series" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "thumbnail" varchar
);

CREATE TABLE "movies" (
  "id" integer PRIMARY KEY,
  "title" varchar,
  "duration" integer,
  "description" varchar,
  "actor_avatars" varchar[],
  "trailer" varchar,
  "file_path" varchar,
  "thumbnail" varchar,
  "series_id" integer,
  "views" integer DEFAULT 0,
  "stars" integer DEFAULT 0,
  "created_at" timestamp DEFAULT 'now()'
);

CREATE TABLE "movie_genres" (
  "movie_id" integer,
  "genre_id" integer,
  "primary" key(movie_id,genre_id)
);

CREATE TABLE "rooms" (
  "id" integer PRIMARY KEY,
  "movie_id" integer,
  "created_by" varchar,
  "created_at" timestamp DEFAULT 'now()'
);

CREATE TABLE "user_rooms" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "room_id" integer,
  "joined_at" timestamp
);

CREATE TABLE "votes" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "movie_id" integer,
  "stars" integer
);

CREATE TABLE "user_rooms_chat" (
  "id" integer PRIMARY KEY,
  "room_id" integer,
  "user_id" integer,
  "message" varchar,
  "sent_at" timestamp DEFAULT 'now()'
);

COMMENT ON COLUMN "series"."thumbnail" IS 'Path to the thumbnail image';

COMMENT ON COLUMN "movies"."file_path" IS 'Path to the movie file';

COMMENT ON COLUMN "movies"."thumbnail" IS 'Path to the thumbnail image';

COMMENT ON COLUMN "movies"."series_id" IS 'ID of the series the movie belongs to, can be null';

COMMENT ON COLUMN "movies"."views" IS 'Number of views';

COMMENT ON COLUMN "movies"."stars" IS 'Total star votes';

COMMENT ON COLUMN "votes"."stars" IS 'Star vote of the user for the movie';

ALTER TABLE "user_rooms_chat" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_rooms_chat" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "rooms" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");

ALTER TABLE "user_rooms" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_rooms" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "movie_genres" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");

ALTER TABLE "movie_genres" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("id");

ALTER TABLE "movies" ADD FOREIGN KEY ("series_id") REFERENCES "series" ("id");

ALTER TABLE "votes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "votes" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");

ALTER TABLE "movies" ADD FOREIGN KEY ("id") REFERENCES "movies" ("stars");
