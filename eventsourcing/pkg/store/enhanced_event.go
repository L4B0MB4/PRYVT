package store

import (
	"time"

	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/models"
	"github.com/google/uuid"
)

type enhancedEvent struct {
	models.Event
	timestamp time.Time
	id        uuid.UUID
}
