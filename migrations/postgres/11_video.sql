-- +migrate Up
CREATE TABLE video
(
    id SERIAL PRIMARY KEY,
    adgroup_id integer  REFERENCES adgroup (id),
    folder_id integer  REFERENCES folder (id),
    creative_id text,
    filepath text,
    youtube_id text,
    title text,
    description text,
    created_at timestamp with time zone NOT NULL DEFAULT now()

);
-- +migrate Down
DROP TABLE video;