-- name: CreatePaste :exec
INSERT INTO
   pastes ( id, content, expires_at, visibility, language, password, author_name ) 
VALUES
   (
      $1, $2, $3, $4, $5, $6, $7 
   )
;
