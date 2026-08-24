-- +goose Up
CREATE TABLE orders (
    id      SERIAL      PRIMARY KEY,
    user_id INTEGER     NOT NULL,
    status  VARCHAR(16) NOT NULL
);

-- +goose Down
DROP TABLE orders;
