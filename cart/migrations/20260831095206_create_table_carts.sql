-- +goose Up
CREATE TABLE carts(
    id      SERIAL  PRIMARY KEY,
    user_id INTEGER NOT NULL,
    sku     INTEGER NOT NULL,
    count   INTEGER NOT NULL
);

-- +goose Down
DROP TABLE carts;
