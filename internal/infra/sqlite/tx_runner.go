package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/tanaxer01/pedalea/internal/core/rental"
)

type TxRunner struct {
	db *sqlx.DB
}

func NewTxRunner(db *sqlx.DB) *TxRunner {
	return &TxRunner{db: db}
}

func (r *TxRunner) WithinTx(fn func(bikeRepo rental.BikeRepo, rentalRepo rental.RentalRepo) error) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	bikeRepo := NewBikeRepository(tx)
	rentalRepo := NewRentalRepository(tx)

	if err := fn(bikeRepo, rentalRepo); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}
