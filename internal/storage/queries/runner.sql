-- name: GetRunner :one
SELECT *
FROM runner
WHERE id = $1;