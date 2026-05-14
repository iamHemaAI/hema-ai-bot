-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    tg_id       INT PRIMARY KEY,
    tg_username TEXT NOT NULL,
    name        TEXT
);

-- +goose Down
DROP TABLE IF EXISTS users;
