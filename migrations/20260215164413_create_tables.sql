-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	hashed_password TEXT NOT NULL,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bikes (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	is_available BOOLEAN NOT NULL DEFAULT 1,
	latitude NUMERIC NOT NULL,
	longitude NUMERIC NOT NULL,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS rentals (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	bike_id INTEGER NOT NULL,
	status TEXT NOT NULL DEFAULT 'running',
	start_time INTEGER NOT NULL,
	end_time INTEGER,
	start_latitude NUMERIC NOT NULL,
	start_longitude NUMERIC NOT NULL,
	end_latitude NUMERIC,
	end_longitude NUMERIC,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (bike_id) REFERENCES bikes(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE bikes;
DROP TABLE rentals;
-- +goose StatementEnd
