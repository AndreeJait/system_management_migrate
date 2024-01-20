
-- +migrate Up
CREATE TABLE IF NOT EXISTS info_savings (
    id bigserial PRIMARY KEY,
    name varchar default 'primary_saving',
    value bigint not null,
    created_at TIMESTAMP(6) default current_timestamp,
    updated_at TIMESTAMP(6) default current_timestamp,
    user_id bigserial not null,
    constraint fk_info_saving_users FOREIGN KEY (user_id) REFERENCES users (id)
);
-- +migrate Down
DROP TABLE IF EXISTS info_savings;