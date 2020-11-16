-- +migrate Up
alter table folder add column rules_mask text;
-- +migrate Down