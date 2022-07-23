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
