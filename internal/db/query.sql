-- name: GetAccountByEmail :one
SELECT * FROM accounts
WHERE email=$1;

-- name: GetAccountById :one
SELECT * FROM accounts
WHERE id=$1;

-- name: IsExistByEmail :one
SELECT EXISTS (
    SELECT 1 FROM accounts WHERE email=$1
);

-- name: CreateAccount :one
INSERT INTO accounts (id, email, password)
VALUES ($1, $2, $3)
RETURNING *;
