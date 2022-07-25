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

-- name: GetCurrenciesByCode :many
SELECT * FROM currency
WHERE code = $1 AND updated_at BETWEEN SYMMETRIC $2 AND $3;

-- name: GetAllCurrencies :many
SELECT * FROM currency
WHERE code != 'ALL' AND updated_at BETWEEN SYMMETRIC $1 AND $2;


-- name: CreateLog :one
INSERT INTO log (
  duration, code, request, created_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;