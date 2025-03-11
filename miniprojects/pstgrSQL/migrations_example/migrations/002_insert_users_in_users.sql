-- +goose Up
INSERT INTO users(name, email)
VALUES('Daniel', 'daniel@mail.com'),
      ('Vlad', 'vla@mail.com'),
      ('Alex', 'sd@mail.com'),
      ('Sonya', 'fdf@mail.com');

-- +goose Down
DELETE FROM users;