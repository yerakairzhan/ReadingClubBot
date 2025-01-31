// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: admins.sql

package db

import (
	"context"
)

const getAdminByID = `-- name: GetAdminByID :one
SELECT user_id
FROM admins
WHERE user_id = $1
`

func (q *Queries) GetAdminByID(ctx context.Context, userID string) (string, error) {
	row := q.db.QueryRowContext(ctx, getAdminByID, userID)
	var user_id string
	err := row.Scan(&user_id)
	return user_id, err
}
