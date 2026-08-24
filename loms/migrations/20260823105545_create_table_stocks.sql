-- +goose Up
CREATE TABLE stocks (
    sku         INTEGER PRIMARY KEY,
    count       INTEGER NOT NULL,
    reserved    INTEGER NOT NULL
);

-- +goose Down
DROP TABLE stocks;
