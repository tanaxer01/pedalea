package sqlite

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type RentalRepository struct {
	db *sqlx.DB
}

func NewRentalRepository(db *sqlx.DB) *RentalRepository {
	return &RentalRepository{db: db}
}

func (r *RentalRepository) InsertRental(data pedalea.RentalData) error {
	_, err := r.db.NamedExec(
		`INSERT INTO rentals (
			user_id, bike_id, start_time, start_latitude, start_longitude, duration, cost
		) VALUES (
			:user_id, :bike_id, :start_time, :start_latitude, :start_longitude, :duration, :cost
		)`, data)

	if isDuplicated(err) {
		return err
	} else if err != nil {
		return err
	}

	return err
}

func (r *RentalRepository) UpdateRental(ID int, data pedalea.RentalData) error {
	res, err := r.db.Exec(
		`UPDATE rentals
		 SET status = COALESCE($1, status),
				 start_time = COALESCE($2, start_time),
				 end_time = COALESCE($3, end_time),
				 start_latitude = COALESCE($4, start_latitude),
				 start_longitude = COALESCE($5, start_longitude),
				 end_latitude = COALESCE($6, end_latitude),
				 end_longitude = COALESCE($7, end_longitude),
				 duration = COALESCE($8, duration),
				 cost = COALESCE($9, cost),
				 updated_at = COALESCE($10, updated_at)
			WHERE id = $11`,
		data.Status,
		data.StartTime,
		data.EndTime,
		data.StartLatitude,
		data.StartLongitude,
		data.EndLatitude,
		data.EndLongitude,
		data.Duration,
		data.Cost,
		time.Now(),
		ID,
	)

	// TODO: Same as usr
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

func (r *RentalRepository) GetRentalByID(ID int) (*pedalea.Rental, error) {
	rental := pedalea.Rental{}
	err := r.db.Get(&rental, "SELECT  id, user_id, bike_id, status, start_time, end_time, start_latitude, start_longitude, end_latitude, end_longitude, duration, cost, created_at, updated_at FROM rentals WHERE id = $1", ID)

	if isNotFound(err) {
		return nil, pedalea.ErrRentalNotFound
	} else if err != nil {
		return nil, err
	}

	return &rental, nil
}

func (r *RentalRepository) GetRentalByUserID(ID int) (*pedalea.Rental, error) {
	rental := pedalea.Rental{}
	err := r.db.Get(&rental, "SELECT  id, user_id, bike_id, status, start_time, end_time, start_latitude, start_longitude, end_latitude, end_longitude, duration, cost, created_at, updated_at FROM rentals WHERE user_id = $1 AND status = 'running'", ID)

	if isNotFound(err) {
		return nil, pedalea.ErrRentalNotFound
	} else if err != nil {
		return nil, err
	}

	return &rental, nil
}

func (r *RentalRepository) ListRentals() ([]pedalea.Rental, error) {
	var rentals []pedalea.Rental
	err := r.db.Select(&rentals, "SELECT id, user_id, bike_id, status, start_time, end_time, start_latitude, start_longitude, end_latitude, end_longitude, duration, cost, created_at, updated_at FROM rentals")

	if err != nil {
		return nil, err
	}

	if rentals == nil {
		rentals = []pedalea.Rental{}
	}

	return rentals, nil
}

func (r *RentalRepository) ListUserRentals(userID int) ([]pedalea.Rental, error) {
	var rentals []pedalea.Rental
	err := r.db.Select(&rentals, "SELECT id, user_id, bike_id, status, start_time, end_time, start_latitude, start_longitude, end_latitude, end_longitude, duration, cost, created_at, updated_at FROM rentals WHERE user_id = $1", userID)

	if err != nil {
		return nil, err
	}

	if rentals == nil {
		rentals = []pedalea.Rental{}
	}

	return rentals, nil
}

func (r *RentalRepository) ListOverlappingRentals(userID, bikeID int) ([]pedalea.Rental, error) {
	var rentals []pedalea.Rental
	err := r.db.Select(&rentals, "SELECT id, user_id, bike_id, status, start_time, end_time, start_latitude, start_longitude, end_latitude, end_longitude, duration, cost, created_at, updated_at FROM rentals WHERE  status = 'running' AND (user_id = $1 OR bike_id = $2)", userID, bikeID)

	if err != nil {
		return nil, err
	}

	if rentals == nil {
		rentals = []pedalea.Rental{}
	}

	return rentals, nil
}
