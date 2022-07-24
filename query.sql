-- name: ListCurrencies :many
SELECT * FROM currency
ORDER BY code;

-- name: CreateCurrency :one
INSERT INTO currency (
  code, value, updated_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreateLog :one
INSERT INTO log (
  duration, code, request, created_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;