
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY ,
    full_name varchar not null,
    username varchar not null unique,
    is_active boolean default TRUE,
    password varchar,
    role varchar not null,
    created_at TIMESTAMP(6) default current_timestamp,
    updated_at TIMESTAMP(6) default current_timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS users;