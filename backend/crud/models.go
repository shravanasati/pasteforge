// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package crud

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Paste struct {
	ID         string
	Content    string
	CreatedAt  pgtype.Timestamp
	ExpiresAt  pgtype.Timestamp
	Visibility string
	Language   string
	Password   pgtype.Text
}

type User struct {
	Username string
	Sessions []byte
}
