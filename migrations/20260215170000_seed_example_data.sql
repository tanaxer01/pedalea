-- +goose Up
-- +goose StatementBegin
INSERT INTO users (email, hashed_password, first_name, last_name)
VALUES
  ('alice@example.com', '$2a$10$examplehashforalice', 'Alice', 'Rider'),
  ('bob@example.com', '$2a$10$examplehashforbob', 'Bob', 'Pedal'),
  ('carol@example.com', '$2a$10$examplehashforcarol', 'Carol', 'Chain');

INSERT INTO bikes (is_available, price_per_minute, latitude, longitude)
VALUES
  (1, 2, 40.712776, -74.005974),
  (0, 3, 40.713776, -74.004974),
  (1, 1, 40.714776, -74.003974),
  (1, 2, 40.715776, -74.002974);

INSERT INTO rentals (
  user_id, bike_id, status, start_time, end_time,
  start_latitude, start_longitude, end_latitude, end_longitude,
  duration, cost
)
VALUES
  ((SELECT id FROM users WHERE email = 'alice@example.com'), (SELECT id FROM bikes WHERE latitude = 40.713776 AND longitude = -74.004974), 'running', 1700000000, NULL, 40.713776, -74.004974, NULL, NULL, 0, 0),
  ((SELECT id FROM users WHERE email = 'bob@example.com'), (SELECT id FROM bikes WHERE latitude = 40.714776 AND longitude = -74.003974), 'ended', 1699990000, 1699993600, 40.714776, -74.003974, 40.715776, -74.002974, 60, 60);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM rentals
WHERE user_id IN (SELECT id FROM users WHERE email IN ('alice@example.com', 'bob@example.com', 'carol@example.com'));

DELETE FROM bikes
WHERE latitude IN (40.712776, 40.713776, 40.714776, 40.715776)
  AND longitude IN (-74.005974, -74.004974, -74.003974, -74.002974);

DELETE FROM users WHERE email IN ('alice@example.com', 'bob@example.com', 'carol@example.com');
-- +goose StatementEnd
