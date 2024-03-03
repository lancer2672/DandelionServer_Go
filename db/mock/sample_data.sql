-- Users
INSERT INTO users (id, username) VALUES (1, 'user1');
INSERT INTO users (id, username) VALUES (2, 'user2');

-- Genres
INSERT INTO genres (id, name) VALUES (1, 'Action');
INSERT INTO genres (id, name) VALUES (2, 'Comedy');

-- Series
INSERT INTO series (id, name, thumbnail) VALUES (1, 'Series1', 'path/to/series1.png');
INSERT INTO series (id, name, thumbnail) VALUES (2, 'Series2', 'path/to/series2.png');

-- Movies
INSERT INTO movies (id, title, duration, description, actor_avatars, trailer, file_path, thumbnail, series_id, views, stars, created_at) VALUES (1, 'Elephant Dream', 600, 'The first Blender Open Movie from 2006', ARRAY['https://picsum.photos/200', 'https://picsum.photos/200'], 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4', 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/images/ElephantsDream.jpg', 1, 100, 5, '2024-03-01 00:00:00');
INSERT INTO movies (id, title, duration, description, actor_avatars, trailer, file_path, thumbnail, series_id, views, stars, created_at) VALUES (2, 'Big Buck Bunny', 600, 'Big Buck Bunny tells the story of a giant rabbit with a heart bigger than himself. When one sunny day three rodents rudely harass him, something snaps... and the rabbit aint no bunny anymore! In the typical cartoon tradition he prepares the nasty rodents a comical revenge', ARRAY['https://picsum.photos/200', 'https://picsum.photos/200'], 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4', 'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/images/BigBuckBunny.jpg', 2, 200, 4, '2024-03-02 00:00:00');

-- Movie Genres
INSERT INTO movie_genres (movie_id, genre_id) VALUES (1, 1);
INSERT INTO movie_genres (movie_id, genre_id) VALUES (1, 2);
INSERT INTO movie_genres (movie_id, genre_id) VALUES (2, 1);

-- Rooms
INSERT INTO rooms (id, movie_id, current_time, is_playing, created_at) VALUES (1, 1, 0, true, '2024-03-01 00:00:00');
INSERT INTO rooms (id, movie_id, current_time, is_playing, created_at) VALUES (2, 2, 0, false, '2024-03-02 00:00:00');

-- User Rooms
INSERT INTO user_rooms (id, user_id, room_id, joined_at) VALUES (1, 1, 1, '2024-03-01 00:00:00');
INSERT INTO user_rooms (id, user_id, room_id, joined_at) VALUES (2, 2, 2, '2024-03-02 00:00:00');

-- Votes
INSERT INTO votes (id, user_id, movie_id, stars) VALUES (1, 1, 1, 5);
INSERT INTO votes (id, user_id, movie_id, stars) VALUES (2, 2, 2, 4);
