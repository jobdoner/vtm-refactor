-- +migrate Up
ALTER TABLE matching_folders ADD COLUMN deleted_at timestamp with time zone;
-- +migrate Down
