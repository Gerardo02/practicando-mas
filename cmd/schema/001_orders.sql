-- +goose Up
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    users VARCHAR NOT NULL,
    total VARCHAR NOT NULL,
    discount VARCHAR NOT NULL,
    payment VARCHAR NOT NULL,
    shipping VARCHAR NOT NULL
);
-- +goose Down
DROP TABLE orders;
