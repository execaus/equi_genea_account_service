-- name: GetAccountFromEmail :one
SELECT * FROM accounts
WHERE email=$1;

-- name: GetAccountFromId :one
SELECT * FROM accounts
WHERE id=$1;