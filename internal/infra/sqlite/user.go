package sqlite

import (
	"database/sql"

	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) InsertUser(user *pedalea.InsertUser) error {
	return nil
}

func (r *UserRepository) UpdateUser(ID int, data pedalea.UserData) error {
	return nil
}

func (r *UserRepository) GetUserByID(ID int) (*pedalea.User, error) {
	return nil, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*pedalea.User, error) {
	return nil, nil
}
