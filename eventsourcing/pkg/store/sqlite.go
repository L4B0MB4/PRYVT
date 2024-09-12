package store

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

type DatabaseConnection struct {
	initialized bool
	db          *sql.DB
}

func (d *DatabaseConnection) SetUp() {
	db, err := sql.Open("sqlite3", "./eventstore.db")
	if err != nil {

		log.Info().Err(err).Msg("Opening sqlite connection")
		return
	}
	if createEventTable(db) != nil {
		return
	}
	if createAggregateStateTable(db) != nil {
		return
	}
	if createAggregateSnapshotTable(db) != nil {
		return
	}
	d.db = db
	d.initialized = true
}

func createEventTable(db *sql.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS events (id TEXT PRIMARY KEY, timestamp INTEGER,Name TEXT, version_0 INTEGER,version_1 INTEGER,data BLOB,UNIQUE(version_0, version_1) ON CONFLICT FAIL)")
	if err != nil {

		log.Info().Err(err).Msg("Preparing statement for events table")
		return err
	}
	_, err = stmt.Exec()
	if err != nil {

		log.Info().Err(err).Msg("Creating events table")
		return err
	}
	return nil
}

func createAggregateStateTable(db *sql.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS aggregate_state (id TEXT PRIMARY KEY, name TEXT,version_0 INTEGER,version_1 INTEGER,UNIQUE(version_0, version_1) ON CONFLICT FAIL )")
	if err != nil {

		log.Info().Err(err).Msg("Preparing statement for aggregate_state table")
		return err
	}
	_, err = stmt.Exec()
	if err != nil {

		log.Info().Err(err).Msg("Creating aggregate_state table")
		return err
	}
	return nil
}

func createAggregateSnapshotTable(db *sql.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS aggregate_snapshots (id TEXT PRIMARY KEY, name TEXT,version_0 INTEGER,version_1 INTEGER,UNIQUE(version_0, version_1) ON CONFLICT FAIL )")
	if err != nil {

		log.Info().Err(err).Msg("Preparing statement for aggregate_snapshots table")
		return err
	}
	_, err = stmt.Exec()
	if err != nil {

		log.Info().Err(err).Msg("Creating aggregate_snapshots table")
		return err
	}
	return nil
}

func (d *DatabaseConnection) GetDbConnection() (*sql.DB, error) {
	if !d.initialized {
		return nil, errors.New("DatabaseConnection not properly initialized")
	}
	return d.db, nil
}
