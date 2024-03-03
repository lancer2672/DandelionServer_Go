-- name: GetRoom :one
SELECT * FROM rooms
WHERE id = $1 LIMIT 1;

-- name: GetRoomsByUser :many
SELECT * FROM rooms
JOIN user_rooms ON rooms.id = user_rooms.room_id
WHERE user_rooms.user_id = $1
ORDER BY rooms.id;
