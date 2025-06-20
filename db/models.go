// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Post struct {
	ID        int64
	ThreadID  pgtype.Int4
	AuthorID  pgtype.Int4
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Thread struct {
	ID        int64
	Title     string
	AuthorID  pgtype.Int4
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
