-- +migrate Up
CREATE TABLE folder
(
    id SERIAL PRIMARY KEY,
    folder_type text,
    name text,
    created_at timestamp with time zone  DEFAULT NOW()
);

-- +migrate Down