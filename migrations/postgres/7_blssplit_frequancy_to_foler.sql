-- +migrate Up
ALTER TABLE folder ADD COLUMN frequency TEXT;
ALTER TABLE folder ADD COLUMN bls_split BOOLEAN;
ALTER TABLE campaign DROP COLUMN frequency;
ALTER TABLE adgroup DROP COLUMN bls_split ;
-- +migrate Down
ALTER TABLE campaign ADD COLUMN frequency TEXT;
ALTER TABLE adgroup ADD COLUMN bls_split TEXT;
ALTER TABLE folder DROP COLUMN frequency;
ALTER TABLE folder DROP COLUMN bls_split;