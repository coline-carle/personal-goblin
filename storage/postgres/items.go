package postgres

import (
	"database/sql"

	"github.com/pkg/errors"
)

// GetWatchAuctionsItems return item ids of all items we want to watch at the auction house.g
func (s *StoragePostgres) GetWatchAuctionsItems(tx *sql.Tx) ([]uint64, error) {
	var ret []uint64
	rows, err := tx.Query("SELECT id from items WHERE watch_auctions = TRUE")
	if err != nil {
		return nil, errors.Wrap(err, "internal server error")
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return nil, errors.Wrap(err, "internal server error")
		}
		ret = append(ret, uint64(id))
	}
	return ret, nil
}
