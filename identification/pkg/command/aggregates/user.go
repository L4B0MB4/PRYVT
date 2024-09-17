package aggregates

import (
	"fmt"
	"time"

	"github.com/L4B0MB4/EVTSRC/pkg/client"
	"github.com/L4B0MB4/EVTSRC/pkg/models"
	"github.com/L4B0MB4/PRYVT/identification/pkg/command/events"
)

type UserAggregate struct {
	name       string
	changeDate time.Time
	events     []models.Event
}

func NewUserAggregate() *UserAggregate {
	aggregateId := "idntification.UserAggregate"
	c, err := client.NewEventSourcingHttpClient("http://localhost:5155")
	if err != nil {
		return nil
	}
	iter, err := c.GetEventsOrdered(aggregateId)
	if err != nil {
		return nil
	}
	ua := &UserAggregate{
		events: []models.Event{},
	}

	for {
		ev, ok := iter.Next()
		if !ok {
			break
		}
		ua.addEvent(ev)
	}
	return ua
}

func (ua *UserAggregate) addEvent(ev interface{}) {
	switch v := ev.(type) {
	case *events.NameChangeEvent:
		fmt.Printf("Twice %v ", v)
	default:
		fmt.Errorf("NO KNOWN EVENT %v", ev)
	}
}

func (ua *UserAggregate) ChangeName(name string) {
	if ua.name != name && len(name) <= 50 && ua.changeDate.Sub(time.Now()).Minutes() > 1 {
		ua.addEvent(events.NameChangeEvent{Name: name})
	}
}
