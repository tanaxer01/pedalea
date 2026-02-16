package sqlite

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mattn/go-sqlite3"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) InsertUser(user *pedalea.InsertUser) error {
	_, err := r.db.Exec(
		`INSERT INTO users (email, first_name, last_name, hashed_password) VALUES (?, ?, ?, ?)`,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Password,
	)

	if err == nil {
		return nil
	}

	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return pedalea.ErrUserAlreadyExists
		}
	}

	return err
}

func (r *UserRepository) UpdateUser(ID int, data pedalea.UserData) error {
	res, err := r.db.Exec(
		`UPDATE users SET email = ?, first_name = ?, last_name = ?, updated_at = ? WHERE id = ?`,
		data.Email,
		data.FirstName,
		data.LastName,
		time.Now(),
		ID,
	)

	if err != nil {
		return err
	}

	updated, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if updated == 0 {
		return pedalea.ErrUserNotFound
	}

	return nil
}

func (r *UserRepository) GetUserByID(ID int) (*pedalea.User, error) {
	var user pedalea.User

	err := r.db.
		QueryRow("SELECT id, email, first_name, last_name, hashed_password, created_at, updated_at FROM users WHERE id = ?", ID).
		Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.HashedPassword, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*pedalea.User, error) {
	var user pedalea.User

	err := r.db.
		QueryRow("SELECT id, email, first_name, last_name, hashed_password, created_at, updated_at FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.HashedPassword, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pedalea.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
