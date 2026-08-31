-- name: CreateItem :one
INSERT INTO carts(user_id, sku, count) VALUES ($1, $2, $3) RETURNING *;
