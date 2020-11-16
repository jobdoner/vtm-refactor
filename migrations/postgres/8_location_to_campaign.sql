-- +migrate Up
ALTER TABLE campaign ADD COLUMN location TEXT;

-- +migrate Down
ALTER TABLE campaign DROP COLUMN location;