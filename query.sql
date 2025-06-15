-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetThread :one
SELECT * FROM threads
WHERE id = $1 LIMIT 1;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;