package store

import (
	"database/sql"
	"time"

	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type EventRepository struct {
	store *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	if db == nil {
		return nil
	}
	return &EventRepository{store: db}
}

func (e *EventRepository) AddEvent(event *models.Event) {
	eEvent := &enhancedEvent{
		Event:     *event,
		timestamp: time.Now(),
		id:        uuid.New(),
	}
	e.addEvent(eEvent)
}

func (e *EventRepository) addEvent(event *enhancedEvent) error {
	// Prepare the insert statement
	stmt, err := e.store.Prepare(`
        INSERT INTO events (id, timestamp, Name, version_0, version_1, data)
        VALUES (?, ?, ?, ?, ?, ?)
        ON CONFLICT(version_0, version_1) DO NOTHING
    `)
	if err != nil {
		log.Info().Err(err).Msg("Preparing insert statement for events table")
		return err
	}
	defer stmt.Close()

	// Execute the statement with the provided values
	_, err = stmt.Exec(id, timestamp, name, version0, version1, data)
	if err != nil {
		log.Info().Err(err).Msg("Inserting into events table")
		return err
	}

	log.Info().Msg("Insert successful")
	return nil
}
