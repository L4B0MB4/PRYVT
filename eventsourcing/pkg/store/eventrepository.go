package store

import (
	"database/sql"
	"errors"
	"fmt"
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

	tx, err := e.store.Begin()
	stmt, err := e.store.Prepare(`
        INSERT INTO events (id, aggregateId,timestamp_0 ,timestamp_1, Name, version_0, version_1, data)
        VALUES (?,?,?,?,?,?,?,?)
    `)
	if err != nil {
		log.Info().Err(err).Msg("Preparing insert statement for events table")
		return err
	}
	defer stmt.Close()

	_, err = tx.Stmt(stmt).Exec(event.id, event.AggregateType, t0, t1, event.Name, v0, v1, event.Data)
	if err != nil {
		tx.Rollback()
		log.Warn().Err(err).Msg("ABORTED TX")
		return err
	}

	stmtAgg, err := e.store.Prepare(`
        INSERT INTO aggregate_state(id,version_0, version_1)
        VALUES (?,?,?)
    `)
	if err != nil {
		log.Info().Err(err).Msg("Preparing insert statement for events table")
		return err
	}
	defer stmt.Close()

	_, err = tx.Stmt(stmtAgg).Exec(event.AggregateType, v0, v1)
	if err != nil {
		tx.Rollback()
		log.Warn().Err(err).Msg("ABORTED TX")
		return err
	}
	tx.Commit()

	return nil
}

func (e *EventRepository) GetEventsForAggregate(aggregateType string) ([]models.Event, error) {

	// Prepare the SQL query
	query := `
		SELECT events.Name, events.version_0, events.version_1, events.data 
		FROM events 
		JOIN aggregate_state 
			ON events.aggregateId = aggregate_state.id 
		 	and events.version_0 = aggregate_state.version_0 
			and events.version_1 = aggregate_state.version_1 
		WHERE aggregate_state.id = ?
		ORDER BY events.version_0 ASC, events.version_1 ASC
	`

	stmt, err := e.store.Prepare(query)
	if err != nil {
		log.Warn().Err(err).Msg("Error preparing statement")
		return nil, errors.New("COULD NOT PREPARE STATMENT FOR QUERY EVENTS")
	}
	defer stmt.Close()

	// Execute the query
	rows, err := stmt.Query(aggregateType)
	if err != nil {
		log.Warn().Err(err).Msg("Error running query statement")
		return nil, errors.New("COULD NOT QUERY EVENTS")
	}
	defer rows.Close()

	// Initialize a slice to hold all events
	var events []models.Event

	// Iterate over the results and store them in the slice
	for rows.Next() {
		var event models.Event

		var v0 int32
		var v1 int32
		err = rows.Scan(&event.Name, &v0, &v1, &event.Data)
		if err != nil {
			log.Warn().Err(err).Msg("Error scanning rows")
			return nil, errors.New("COULD NOT RETRIEVE EVENT")
		}

		// Append the event to the slice
		events = append(events, event)
	}

	// Check for any error that might have occurred during iteration
	if err = rows.Err(); err != nil {
		log.Warn().Err(err).Msg("Error checking row errors")
		return nil, errors.New("COULD NOT RETRIEVE ALL EVENTS")
	}

	// Print out the events or handle them as needed
	for _, event := range events {
		fmt.Printf("Event Name: %s, Version: %v \n", event.Name, event.Version)
	}
	return events, nil
}
