
-- +migrate Up
CREATE TABLE IF NOT EXISTS target_used_moneys (
    id bigserial PRIMARY KEY,
    purpose varchar not null,
    value bigint not null,
    month int not null,
    year int not null,
    created_at TIMESTAMP(6) default current_timestamp,
    updated_at TIMESTAMP(6) default current_timestamp,
    user_id bigserial not null,
    constraint fk_info_saving_users FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +migrate Down
DROP TABLE IF EXISTS used_moneys;
