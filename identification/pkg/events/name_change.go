package events

import "github.com/L4B0MB4/EVTSRC/pkg/models"

type NameChangedEvent struct {
	Name string
}

func NewNameChangedEvent(name string) *models.Event {
	b := UnsafeSerializeAny(NameChangedEvent{Name: name})
	return &models.Event{
		Name: "NameChangeEvent",
		Data: b,
	}
}
