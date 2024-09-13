package store

import (
	"database/sql"
	"time"

	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/helper"
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

func (e *EventRepository) AddEvent(event *models.Event) error {
	eEvent := &enhancedEvent{
		Event:     *event,
		timestamp: time.Now(),
		id:        uuid.New(),
	}
	return e.addEvent(eEvent)
}

func (e *EventRepository) addEvent(event *enhancedEvent) error {
	t0, t1, err := helper.SplitVersion(event.timestamp.UnixMicro())
	if err != nil {
		return err
	}

	v0, v1, err := helper.SplitVersion(event.Version)
	if err != nil {
		return err
	}
	stmt, err := e.store.Prepare(`
        INSERT INTO events (id, timestamp_0 ,timestamp_1, Name, version_0, version_1, data)
        VALUES (?, ?,?, ?, ?, ?, ?)
    `)
	if err != nil {
		log.Info().Err(err).Msg("Preparing insert statement for events table")
		return err
	}
	defer stmt.Close()

	// Execute the statement with the provided values
	_, err = stmt.Exec(event.id, t0, t1, event.Name, v0, v1, event.Data)
	if err != nil {
		log.Error().Err(err).Msg("Inserting into events table")
		return err
	}

	log.Info().Msg("Insert successful")
	return nil
}
