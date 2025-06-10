-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;

-- name: GetAllTask :many
SELECT * FROM tasks;