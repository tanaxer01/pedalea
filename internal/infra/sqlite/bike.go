package sqlite

import (
	"database/sql"

	"github.com/tanaxer01/pedalea/pkg/pedalea"
)

type BikeRepository struct {
	db *sql.DB
}

func NewBikeRepository(db *sql.DB) *BikeRepository {
	return &BikeRepository{db: db}
}

func (r *BikeRepository) UpdateBike(bikeID int, data pedalea.BikeData) error {
	return nil
}

func (r *BikeRepository) GetBikeByID(bikeID int) (*pedalea.Bike, error) {
	return nil, nil
}

func (r *BikeRepository) ListAvailableBikes() ([]pedalea.Bike, error) {
	var bikes []pedalea.Bike

	rows, err := r.db.Query("SELECT id, is_available, latitude, longitude, created_at, updated_at FROM bikes WHERE is_available = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bike pedalea.Bike
		err = rows.Scan(&bike.ID, &bike.Available, &bike.Latitude, &bike.Longitude, &bike.CreatedAt, &bike.UpdatedAt)

		if err != nil {
			return nil, err
		}

		bikes = append(bikes, bike)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if bikes == nil {
		bikes = []pedalea.Bike{}
	}

	return bikes, nil
}
