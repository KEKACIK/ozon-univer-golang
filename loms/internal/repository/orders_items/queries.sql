-- name: GetByOrderId :many
SELECT * FROM orders_items WHERE order_id=$1;

-- name: CreateOrderItem :one
INSERT INTO orders_items(order_id, sku, count) VALUES ($1, $2, $3) RETURNING *;
