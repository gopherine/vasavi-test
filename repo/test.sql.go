// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: test.sql

package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTest = `-- name: CreateTest :one
INSERT INTO test (name, bio)
VALUES ($1, $2)
RETURNING id, name, bio
`

type CreateTestParams struct {
	Name string
	Bio  pgtype.Text
}

func (q *Queries) CreateTest(ctx context.Context, arg CreateTestParams) (Test, error) {
	row := q.db.QueryRow(ctx, createTest, arg.Name, arg.Bio)
	var i Test
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const deleteTest = `-- name: DeleteTest :exec
DELETE FROM test
WHERE id = $1
`

func (q *Queries) DeleteTest(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteTest, id)
	return err
}

const getTestByID = `-- name: GetTestByID :one
SELECT id, name, bio FROM test
WHERE id = $1
`

func (q *Queries) GetTestByID(ctx context.Context, id int64) (Test, error) {
	row := q.db.QueryRow(ctx, getTestByID, id)
	var i Test
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const listTests = `-- name: ListTests :many
SELECT id, name, bio FROM test
`

func (q *Queries) ListTests(ctx context.Context) ([]Test, error) {
	rows, err := q.db.Query(ctx, listTests)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Test
	for rows.Next() {
		var i Test
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTest = `-- name: UpdateTest :exec
UPDATE test
SET name = $2, bio = $3
WHERE id = $1
`

type UpdateTestParams struct {
	ID   int64
	Name string
	Bio  pgtype.Text
}

func (q *Queries) UpdateTest(ctx context.Context, arg UpdateTestParams) error {
	_, err := q.db.Exec(ctx, updateTest, arg.ID, arg.Name, arg.Bio)
	return err
}
