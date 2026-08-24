-- +goose Up
CREATE TABLE orders_items (
    id          SERIAL  PRIMARY KEY,
    order_id    BIGINT  NOT NULL,
    sku         INTEGER NOT NULL,
    count       INTEGER NOT NULL,

    FOREIGN KEY (order_id) REFERENCES orders(id)
);

-- +goose Down
DROP TABLE orders_items;
