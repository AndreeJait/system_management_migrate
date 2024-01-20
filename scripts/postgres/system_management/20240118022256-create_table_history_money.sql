
-- +migrate Up
CREATE TABLE IF NOT EXISTS history_moneys (
    id bigserial primary key,
    value bigint not null,
    purpose varchar,
    operation varchar not null,
    created_at TIMESTAMP(6) default current_timestamp,
    updated_at TIMESTAMP(6) default current_timestamp,
    user_id bigserial not null,
    constraint fk_info_saving_users FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +migrate Down
DROP TABLE IF EXISTS history_moneys;