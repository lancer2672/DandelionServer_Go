-- name: GetVotesByUser :many
SELECT * FROM votes
WHERE votes.user_id = $1
ORDER BY votes.id;