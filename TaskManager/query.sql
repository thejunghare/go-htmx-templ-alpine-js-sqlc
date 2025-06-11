-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;

-- name: GetAllTask :many
SELECT * FROM tasks;

-- name: Delete :exec
DELETE FROM tasks
WHERE id = $1;

-- name: UpdateStatus :exec
UPDATE tasks SET status = $2
WHERE id = $1;