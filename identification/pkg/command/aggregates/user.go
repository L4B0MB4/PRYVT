package aggregates

import (
	"fmt"
	"time"

	"github.com/L4B0MB4/EVTSRC/pkg/client"
	"github.com/L4B0MB4/EVTSRC/pkg/models"
	"github.com/L4B0MB4/PRYVT/identification/pkg/events"
	"github.com/google/uuid"
)

type UserAggregate struct {
	name          string
	changeDate    time.Time
	events        []models.Event
	aggregateType string
	aggregateId   uuid.UUID
	client        *client.EventSourcingHttpClient
}

func NewUserAggregate(id uuid.UUID) (*UserAggregate, error) {

	c, err := client.NewEventSourcingHttpClient("http://localhost:5515")
	if err != nil {
		panic(err)
	}
	iter, err := c.GetEventsOrdered(id.String())
	if err != nil {
		return nil, fmt.Errorf("COULDN'T RETRIEVE EVENTS ")
	}
	ua := &UserAggregate{
		client:        c,
		events:        []models.Event{},
		aggregateType: "user",
		aggregateId:   id,
	}

	for {
		ev, ok := iter.Next()
		if !ok {
			break
		}
		ua.addEvent(ev)
	}
	return ua, nil
}

func (ua *UserAggregate) apply(e *events.NameChangedEvent) {

}

func (ua *UserAggregate) addEvent(ev *models.Event) {
	switch ev.Name {
	case "NameChangeEvent":
		e := events.UnsafeDeserializeAny[events.NameChangedEvent](ev.Data)
		ua.apply(e)
	default:
		panic(fmt.Errorf("NO KNOWN EVENT %v", ev))
	}
	if ev.Version == 0 {
		v := len(ua.events) + 1 //for validation we need to start at 1
		ev.Version = int64(v)
		ev.AggregateType = ua.aggregateType
		ev.AggregateId = ua.aggregateId.String()
		ua.events = append(ua.events, *ev)
	}
}

func (ua *UserAggregate) saveChanges() error {
	return ua.client.AddEvents(ua.aggregateId.String(), ua.events)
}
func (ua *UserAggregate) ChangeName(name string) error {
	if ua.name != name && len(name) <= 50 && time.Since(ua.changeDate).Seconds() > 10 {
		ua.addEvent(events.NewNameChangedEvent(name))
		err := ua.saveChanges()
		if err != nil {
			return fmt.Errorf("ERROR ")
		}
		return nil
	}
	return fmt.Errorf("VALIDATING USERNAME FAILED")
}
