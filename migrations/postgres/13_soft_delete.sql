-- +migrate Up
ALTER TABLE folder ADD COLUMN deleted_at timestamp with time zone;
-- +migrate Down
