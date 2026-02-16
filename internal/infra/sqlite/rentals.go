package sqlite

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mattn/go-sqlite3"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type RentalRepository struct {
	db *sql.DB
}

func NewRentalRepository(db *sql.DB) *RentalRepository {
	return &RentalRepository{db: db}
}

func (r *RentalRepository) InsertRental(data pedalea.RentalData) error {
	_, err := r.db.Exec(
		`INSERT INTO rentals (user_id, bike_id, start_time, start_latitude, start_longitude) VALUES (?, ?, ?, ?, ?)`,
		data.UserID,
		data.BikeID,
		data.StartTime,
		data.StartLatitude,
		data.StartLongitude,
	)

	if err == nil {
		return nil
	}

	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return pedalea.ErrRentalAlreadyExists
		}
	}

	return err
}

func (r *RentalRepository) UpdateRental(ID int, data pedalea.RentalData) error {
	res, err := r.db.Exec(
		`UPDATE rentals SET status = ?, start_time = ?, end_time = ?, start_latitude = ?, start_longitude = ?, end_latitude = ?, end_longitude = ?, updated_at = ? WHERE id = ?`,
		data.Status,
		data.StartTime,
		data.EndTime,
		data.StartLatitude,
		data.StartLongitude,
		data.EndLatitude,
		data.EndLongitude,
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
		return pedalea.ErrRentalNotFound
	}

	return err
}

func (r *RentalRepository) ListRentals(userID int) ([]pedalea.Rental, error) {
	var rentals []pedalea.Rental

	rows, err := r.db.Query(
		"SELECT id, user_id, bike_id, status, start_time, end_time, start_latitude, start_longitude, end_latitude, end_longitude FROM rentals",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rental pedalea.Rental
		err = rows.Scan(&rental.ID, &rental.UserID, &rental.BikeID, &rental.Status, &rental.StartTime, &rental.EndTime, &rental.StartLatitude, &rental.StartLongitude, &rental.EndLatitude, &rental.EndLongitude)
		if err != nil {
			return nil, err
		}
		rentals = append(rentals, rental)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if rentals == nil {
		rentals = []pedalea.Rental{}
	}

	return rentals, nil
}

func (r *RentalRepository) ListOverlappingRentals(userID, bikeID int) ([]pedalea.Rental, error) {
	var rentals []pedalea.Rental

	rows, err := r.db.Query(
		"SELECT id, user_id, bike_id, status, start_time, end_time, start_latitude, start_longitude, end_latitude, end_longitude FROM rentals WHERE status = 'running' AND (user_id = ? OR bike_id = ?)",
		userID,
		bikeID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rental pedalea.Rental
		err = rows.Scan(&rental.ID, &rental.UserID, &rental.BikeID, &rental.Status, &rental.StartTime, &rental.EndTime, &rental.StartLatitude, &rental.StartLongitude, &rental.EndLatitude, &rental.EndLongitude)
		if err != nil {
			return nil, err
		}
		rentals = append(rentals, rental)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if rentals == nil {
		rentals = []pedalea.Rental{}
	}

	return rentals, nil
}
