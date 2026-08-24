-- name: ListStocks :many
SELECT *
FROM stocks;

-- name: GetStocks :one
SELECT *
FROM stocks
WHERE sku=$1;

-- name: UpdateStock :exec
UPDATE stocks
SET count=$2,
    reserved=$3
WHERE sku=$1;
