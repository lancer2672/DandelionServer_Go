
-- Thêm dữ liệu cho bảng "users"
INSERT INTO "users" ("id", "username") VALUES
(1, 'user1'),
(2, 'user2'),
(3, 'user3');

-- Thêm dữ liệu cho bảng "genres"
INSERT INTO "genres" ("id", "name") VALUES
(1, 'Action'),
(2, 'Comedy'),
(3, 'Drama');

-- Thêm dữ liệu cho bảng "series"
INSERT INTO "series" ("id", "name", "thumbnail") VALUES
(1, 'Series 1', 'path/to/thumbnail1.jpg'),
(2, 'Series 2', 'path/to/thumbnail2.jpg');

-- Thêm dữ liệu cho bảng "movies"
INSERT INTO "movies" ("id", "title", "duration", "description", "actor_avatars", "trailer", "file_path", "thumbnail", "views", "stars", "created_at") VALUES
(1, 'Movie 1', 120, 'The genus Pisum contains two species, Pisum fulvum and Pisum sativum. The latter is divided into three subspecies: the domesticated pea ssp. sativum, and three wild taxa ssp. elatius and ssp. humile; ssp. humile can be further divided into two varieties, var. humile and var. syriacum. The latter taxon is fully interfertile with the cultigens and can be considered the wild progenitor of pea. For breeding purposes, both ssp. elatius and ssp. humile can be viewed', ARRAY['https://picsum.photos/200', 'https://picsum.photos/200'], 'path/to/trailer1.mp4', 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4', 'path/to/thumbnail1.jpg',  1000, 4, NOW()),
(2, 'Movie 2', 90, 'The genus Pisum contains two species, Pisum fulvum and Pisum sativum. The latter is divided into three subspecies: the domesticated pea ssp. sativum, and three wild taxa ssp. elatius and ssp. humile; ssp. humile can be further divided into two varieties, var. humile and var. syriacum. The latter taxon is fully interfertile with the cultigens and can be considered the wild progenitor of pea. For breeding purposes, both ssp. elatius and ssp. humile can be viewed', ARRAY['https://picsum.photos/200', 'https://picsum.photos/200'], 'path/to/trailer2.mp4', 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4', 'path/to/thumbnail2.jpg',  500, 3, NOW());

-- Thêm dữ liệu cho bảng "movie_genres"
INSERT INTO "movie_genres" ("movie_id", "genre_id") VALUES
(1, 1),
(1, 2),
(2, 2),
(2, 3);

-- Thêm dữ liệu cho bảng "movie_history"
INSERT INTO "movie_history" ("id", "user_id", "movie_id", "watched_duration", "last_watched") VALUES
(1, 1, 1, 60, '2022-01-01 00:00:00'),
(2, 1, 2, 45, '2022-01-02 00:00:00'),
(3, 2, 1, 120, '2022-01-03 00:00:00'),
(4, 3, 2, 90, '2022-01-04 00:00:00');

-- Thêm dữ liệu cho bảng "movies_series"
INSERT INTO "movies_series" ("id", "movie_id", "series_id") VALUES
(1, 1, 1),
(2, 2, 1),
(3, 1, 2),
(4, 2, 2);

-- Thêm dữ liệu cho bảng "votes"
INSERT INTO "votes" ("id", "user_id", "movie_id", "stars") VALUES
(1, 1, 1, 5),
(2, 1, 2, 4),
(3, 2, 1, 5),
(4, 3, 2, 4);

