-- name: CreatePaste :exec
INSERT INTO
   pastes ( id, content, expires_at, visibility, language, password ) 
VALUES
   (
      $1, $2, $3, $4, $5, $6
   )
;
