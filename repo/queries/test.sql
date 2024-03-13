-- name: CreateTest :one
INSERT INTO test (name, bio)
VALUES ($1, $2)
RETURNING id, name, bio;

-- name: GetTestByID :one
SELECT id, name, bio FROM test
WHERE id = $1;

-- name: ListTests :many
SELECT id, name, bio FROM test;

-- name: UpdateTest :exec
UPDATE test
SET name = $2, bio = $3
WHERE id = $1;

-- name: DeleteTest :exec
DELETE FROM test
WHERE id = $1;