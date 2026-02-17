package sqlite

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) InsertUser(user *pedalea.InsertUser) error {
	_, err := r.db.NamedExec(
		`INSERT INTO users (email, first_name, last_name, hashed_password) VALUES (:email, :first_name, :last_name, :password)`,
		user,
	)

	if isDuplicated(err) {
		return err
	} else if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(ID int, data pedalea.UserData) error {
	res, err := r.db.Exec(
		`UPDATE  users SET email = COALLECE($1, email), first_name = COALLECE($2, first_name), last_name = COALLECE($3, last_name), updated_at = $4 WHERE id = $5`,
		data.Email,
		data.FirstName,
		data.LastName,
		time.Now(),
		ID,
	)

	// TODO: Handle better cases where nothing was updated
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
	user := pedalea.User{}
	err := r.db.Get(&user, "SELECT id, email, first_name, last_name, hashed_password, created_at, updated_at FROM users WHERE id = $2", ID)

	if isNotFound(err) {
		return nil, pedalea.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*pedalea.User, error) {
	user := pedalea.User{}
	err := r.db.Get(&user, "SELECT id, email, first_name, last_name, hashed_password, created_at, updated_at FROM users WHERE email = $1", email)

	if isNotFound(err) {
		return nil, pedalea.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) ListUsers() ([]pedalea.User, error) {
	var users []pedalea.User
	err := r.db.Select(&users, "SELECT id, email, first_name, last_name, hashed_password, created_at, updated_at FROM users")

	if err != nil {
		return nil, err
	}

	return users, nil
}
