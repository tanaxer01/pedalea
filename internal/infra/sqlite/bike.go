package sqlite

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type BikeRepository struct {
	db *sqlx.DB
}

func NewBikeRepository(db *sqlx.DB) *BikeRepository {
	return &BikeRepository{db: db}
}

func (r *BikeRepository) InsertBike(bike pedalea.BikeData) error {
	_, err := r.db.NamedExec(`INSERT INTO bikes (latitude, longitude, price_per_minute) VALUES (:latitude, :longitude, :price_per_minute)`, bike)

	if isDuplicated(err) {
		return pedalea.ErrBikeAlreadyExists
	} else if err != nil {
		return err
	}

	return nil
}

func (r *BikeRepository) UpdateBike(ID int, data pedalea.BikeData) error {
	res, err := r.db.Exec(
		`UPDATE bikes SET is_available = COALESCE($1, is_available), latitude = COALESCE($2, latitude), longitude = COALESCE($3, longitude), price_per_minute = COALESCE($4, price_per_minute), updated_at = COALESCE($5, updated_at) WHERE id = $6`,
		data.Available,
		data.Latitude,
		data.Longitude,
		data.PricePerMinute,
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
		return pedalea.ErrBikeNotFound
	}

	return nil
}

func (r *BikeRepository) GetBikeByID(ID int) (*pedalea.Bike, error) {
	var bike pedalea.Bike
	err := r.db.Get(&bike, "SELECT id, is_available, latitude, longitude, price_per_minute, created_at, updated_at FROM bikes WHERE id = $1", ID)

	if isNotFound(err) {
		return nil, pedalea.ErrBikeNotFound
	} else if err != nil {
		return nil, err
	}

	return &bike, nil
}

func (r *BikeRepository) ListAvailableBikes() ([]pedalea.Bike, error) {
	var bikes []pedalea.Bike
	err := r.db.Select(&bikes, "SELECT id, is_available, latitude, longitude, price_per_minute, created_at, updated_at FROM bikes WHERE is_available = true")

	if err != nil {
		return nil, err
	}

	if bikes == nil {
		bikes = []pedalea.Bike{}
	}

	return bikes, nil
}

func (r *BikeRepository) ListBikes() ([]pedalea.Bike, error) {
	var bikes []pedalea.Bike
	err := r.db.Select(&bikes, "SELECT id, is_available, latitude, longitude, price_per_minute, created_at, updated_at FROM bikes")

	if err != nil {
		return nil, err
	}

	if bikes == nil {
		bikes = []pedalea.Bike{}
	}

	return bikes, nil
}
