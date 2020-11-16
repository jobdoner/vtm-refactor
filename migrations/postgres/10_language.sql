-- +migrate Up
ALTER TABLE  adgroup ADD COLUMN languages TEXT[];
ALTER TABLE  adgroup ADD COLUMN ad_schedule TEXT[];

-- +migrate Down
ALTER TABLE adgroup DROP COLUMN languages;
ALTER TABLE  adgroup DROP COLUMN ad_schedule ;