package postgres

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

// Tx wrapper function for database transactions
func (s *StoragePostgres) Tx(fn func(*sql.Tx) error) error {
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "begin transaction")
	}
	err = fn(tx)
	if err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			log.Printf("Transaction rollback failed: %+v", err2)
		}
		return errors.WithStack(err)
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "commit transaction")
	}

	return nil
}
