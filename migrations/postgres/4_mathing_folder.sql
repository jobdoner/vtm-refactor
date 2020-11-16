-- +migrate Up
CREATE TABLE matching_folders
(
    id SERIAL PRIMARY KEY,
    video_folder_id integer REFERENCES folder(id),
    audience_folder_id integer REFERENCES folder(id),
    name text,
    created_at timestamp with time zone DEFAULT NOW()
);

-- +migrate Down
drop table matching_folders;