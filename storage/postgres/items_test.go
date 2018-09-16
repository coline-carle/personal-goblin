package postgres

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetWatchAuctionsItems(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := newTest(db)

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(123).
		AddRow(345).
		AddRow(678)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id from items WHERE watch_auctions = TRUE").
		WillReturnRows(rows)
	mock.ExpectCommit()

	var ids []uint64
	err = storage.Tx(func(tx *sql.Tx) error {
		var err2 error
		ids, err2 = storage.GetWatchAuctionsItems(tx)
		return err2
	})

	if err != nil {
		t.Errorf("got an error during GetWatchAuctionItems transaction: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	expected := []uint64{123, 345, 678}
	assert.Equal(t, expected, ids)
}
