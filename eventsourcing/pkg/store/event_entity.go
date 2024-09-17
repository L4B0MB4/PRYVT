package store

import (
	"time"

	"github.com/L4B0MB4/PRYVT/eventsourcing/pkg/models"
	"github.com/google/uuid"
)

type eventEntity struct {
	models.Event
	timestamp time.Time
	id        uuid.UUID
}
