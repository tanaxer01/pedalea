package sqlite

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type DBRunner interface {
	Exec(query string, args ...any) (sql.Result, error)
	NamedExec(query string, arg any) (sql.Result, error)
	Get(dest any, query string, args ...any) error
	Select(dest any, query string, args ...any) error
}

func NewDB(file string) (*sqlx.DB, error) {
	return sqlx.Connect("sqlite3", file)
}

func isNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func isDuplicated(err error) bool {
	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
		return true
	}

	return false
}
