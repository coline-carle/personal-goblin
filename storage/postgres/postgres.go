package postgres

import (
	"database/sql"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/pkg/errors"

	// database drivers
	_ "github.com/lib/pq"

	// migration  source driver
	_ "github.com/golang-migrate/migrate/source/file"
)

// StoragePostgres Implements storage.Storage
type StoragePostgres struct {
	postgresURL string
	db          *sql.DB
}

// create a new postgres structure for test
func newTest(db *sql.DB) *StoragePostgres {
	return &StoragePostgres{
		db: db,
	}
}

// New create a new StoragePostgres struct
func New(postgresURL string) (*StoragePostgres, error) {
	return &StoragePostgres{
		postgresURL: postgresURL,
	}, nil
}

// Connect connect to the postgres database
func (s *StoragePostgres) Connect() error {
	var err error
	s.db, err = sql.Open("postgres", s.postgresURL)
	if err != nil {
		return errors.Wrap(err, "error connecting to postgres database")
	}

	return nil
}

// ApplyMigrations apply migration to the postgres database
func (s *StoragePostgres) ApplyMigrations() error {
	driver, err := postgres.WithInstance(s.db, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "can't bind migration driver to database driver")
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://storage/postgres/migrations",
		"postgres", driver,
	)
	if err != nil {
		return errors.Wrap(err, "error initializing migration object")
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "error migrating to the latest")
	}

	return nil
}
