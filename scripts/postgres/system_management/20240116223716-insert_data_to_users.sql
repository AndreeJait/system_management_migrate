
-- +migrate Up
INSERT INTO users (username, password, full_name, role) VALUES
        ('panjaitanandree@gmail.com', '$2a$04$LiSuUvol8QlO76ePndhH5OzSc6vdpbewyQbFWSdyHgn3q3xTfXhnG', 'Andree Panjaitan', 'admin');
INSERT INTO info_savings (name, value, user_id) VALUES
        ('saldo utama', '10000', 1);
-- +migrate Down
