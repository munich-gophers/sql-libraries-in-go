// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (name, age, birthday, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)
`

type CreateUserParams struct {
	Name      string
	Age       int64
	Birthday  pgtype.Timestamp
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.Name,
		arg.Age,
		arg.Birthday,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, name, age, birthday, created_at, updated_at, deleted_at FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Age,
			&i.Birthday,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, name, age, birthday, created_at, updated_at, deleted_at FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Age,
		&i.Birthday,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET name = $1 WHERE id = $2
`

type UpdateUserParams struct {
	Name string
	ID   int32
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser, arg.Name, arg.ID)
	return err
}
