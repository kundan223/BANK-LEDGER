-- name: CreateAccount :one
INSERT INTO accounts (
    user_id
)
VALUES ($1)
RETURNING *;


-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1;


-- name: GetAccountForUpdate :one
SELECT *
FROM accounts
WHERE id = $1
FOR UPDATE;