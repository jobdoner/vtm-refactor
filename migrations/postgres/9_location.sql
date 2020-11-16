-- +migrate Up
ALTER TABLE  adgroup ADD COLUMN location TEXT[];

-- +migrate Down
ALTER TABLE adgroup DROP COLUMN location;