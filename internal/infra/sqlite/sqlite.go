package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewBikingDB(file string) (*sql.DB, error) {
	return sql.Open("sqlite3", file)
}
