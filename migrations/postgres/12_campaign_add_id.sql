-- +migrate Up
ALTER TABLE  campaign ADD COLUMN folder_id integer references folder(id);
-- +migrate Down
