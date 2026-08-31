-- name: CreateItem :one
INSERT INTO
carts(user_id, sku, count)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetItemBySku :one
SELECT * FROM carts
WHERE user_id=$1 AND sku=$2;

-- name: GetAllItemByUser :many
SELECT * FROM carts
WHERE user_id=$1;

-- name: UpdateItemCount :exec
UPDATE carts
SET count=$2
WHERE id=$1;

-- name: DeleteItemBySku :exec
DELETE FROM carts
WHERE user_id=$1 AND sku=$2;

-- name: DeleteItemsByUserId :exec
DELETE FROM carts
WHERE user_id=$1;
