-- +goose Up
INSERT INTO stocks(sku, count, reserved) VALUES (1076963, 15, 0);
INSERT INTO stocks(sku, count, reserved) VALUES (5678901, 62, 0);
INSERT INTO stocks(sku, count, reserved) VALUES (4200000, 543, 0);
INSERT INTO stocks(sku, count, reserved) VALUES (2500001, 900, 0);
INSERT INTO stocks(sku, count, reserved) VALUES (8765432, 122, 0);
INSERT INTO stocks(sku, count, reserved) VALUES (7000123, 1, 0);

-- +goose Down
DELETE FROM stocks WHERE sku IN (1076963, 5678901, 4200000, 2500001, 8765432, 7000123);
