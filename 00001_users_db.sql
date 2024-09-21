-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username text,
    userid bigint
);
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS users;
