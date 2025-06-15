-- name: GetUserAndReturnId :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetThreadAndReturnId :one
SELECT * FROM threads
WHERE id = $1 LIMIT 1;

-- name: GetPostAndReturnId :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;