package events

import (
	"time"

	"github.com/L4B0MB4/EVTSRC/pkg/models"
)

type NameChangedEvent struct {
	Name       string
	ChangeDate time.Time
}

func NewNameChangedEvent(name string) *models.ChangeTrackedEvent {
	b := UnsafeSerializeAny(NameChangedEvent{Name: name, ChangeDate: time.Now()})
	return &models.ChangeTrackedEvent{
		Event: models.Event{

			Name: "NameChangeEvent",
			Data: b,
		},
	}
}
