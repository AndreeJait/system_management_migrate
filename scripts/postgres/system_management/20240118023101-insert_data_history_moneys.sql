
-- +migrate Up
INSERT INTO history_moneys (value,purpose, operation, user_id) VALUES
    ('10000', 'init money' ,'+', 1);
-- +migrate Down
