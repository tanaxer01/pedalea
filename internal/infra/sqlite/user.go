package sqlite

import (
	"time"

	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type UserRepository struct {
	db DBRunner
}

func NewUserRepository(db DBRunner) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) InsertUser(user *pedalea.InsertUser) error {
	_, err := r.db.NamedExec(
		`INSERT INTO users (email, first_name, last_name, hashed_password) VALUES (:email, :first_name, :last_name, :password)`,
		user,
	)

	if isDuplicated(err) {
		return pedalea.ErrUserAlreadyExists
	} else if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(ID int, data pedalea.UserData) error {
	res, err := r.db.Exec(
		`UPDATE  users SET email = $1, first_name = $2, last_name = $3, updated_at = $4 WHERE id = $5`,
		data.Email,
		data.FirstName,
		data.LastName,
		time.Now(),
		ID,
	)

	if isDuplicated(err) {
		return pedalea.ErrEmailAlreadyExists
	} else if err != nil {
		return err
	}

	updated, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// NOTE: This case should not be possible if the ID comes from a valid JWT,
	// but we'll handle it anyway. (If error is exposed to the client, this could
	// help some kind of attack).
	if updated == 0 {
		return pedalea.ErrUserNotFound
	}

	return nil
}

func (r *UserRepository) GetUserByID(ID int) (*pedalea.User, error) {
	user := pedalea.User{}
	err := r.db.Get(&user, "SELECT id, email, first_name, last_name, hashed_password, created_at, updated_at FROM users WHERE id = $1", ID)

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
