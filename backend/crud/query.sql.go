// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package crud

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPaste = `-- name: CreatePaste :exec
INSERT INTO
   pastes ( id, content, expires_at, visibility, language, password, author_name ) 
VALUES
   (
      $1, $2, $3, $4, $5, $6, $7 
   )
`

type CreatePasteParams struct {
	ID         string
	Content    string
	ExpiresAt  pgtype.Timestamp
	Visibility string
	Language   string
	Password   pgtype.Text
	AuthorName pgtype.Text
}

func (q *Queries) CreatePaste(ctx context.Context, arg CreatePasteParams) error {
	_, err := q.db.Exec(ctx, createPaste,
		arg.ID,
		arg.Content,
		arg.ExpiresAt,
		arg.Visibility,
		arg.Language,
		arg.Password,
		arg.AuthorName,
	)
	return err
}