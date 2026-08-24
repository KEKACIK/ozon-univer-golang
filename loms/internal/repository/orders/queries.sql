-- name: GetByIdOrder :one
SELECT * FROM orders WHERE id=$1;

-- name: CreateOrder :one
INSERT INTO orders(user_id, Status) VALUES ($1, $2) RETURNING *;

-- name: SetStatusOrder :exec
UPDATE orders SET status=$2 WHERE id=$1;
