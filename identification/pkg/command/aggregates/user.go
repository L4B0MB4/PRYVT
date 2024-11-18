package aggregates

import (
	"fmt"
	"strings"
	"time"

	"github.com/L4B0MB4/EVTSRC/pkg/client"
	"github.com/L4B0MB4/EVTSRC/pkg/models"
	m "github.com/L4B0MB4/PRYVT/identification/pkg/command/models"
	"github.com/L4B0MB4/PRYVT/identification/pkg/events"
	"github.com/google/uuid"
)

type UserAggregate struct {
	displayName   string
	name          string
	passwordHash  string
	email         string
	changeDate    time.Time
	events        []models.ChangeTrackedEvent
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
		events:        []models.ChangeTrackedEvent{},
		aggregateType: "user",
		aggregateId:   id,
		changeDate:    time.Date(2000, 0, 0, 0, 0, 0, 0, time.UTC),
	}

	for {
		ev, ok := iter.Next()
		if !ok {
			break
		}
		changeTrackedEv := models.ChangeTrackedEvent{
			Event: *ev,
			IsNew: false,
		}
		ua.addEvent(&changeTrackedEv)
	}
	return ua, nil
}

func (ua *UserAggregate) apply_DisplayNameChangedEvent(e *events.DisplayNameChangedEvent) {
	ua.displayName = e.DisplayName
	ua.changeDate = e.ChangeDate

}
func (ua *UserAggregate) apply_UserCreatedEvent(e *events.UserCreatedEvent) {
	ua.name = e.Name
	ua.displayName = e.Name
	ua.changeDate = e.CreationDate
	ua.passwordHash = e.PasswordHash
	ua.email = e.Email
}

func (ua *UserAggregate) addEvent(ev *models.ChangeTrackedEvent) {
	switch ev.Name {
	case "NameChangeEvent":
		e := events.UnsafeDeserializeAny[events.DisplayNameChangedEvent](ev.Data)
		ua.apply_DisplayNameChangedEvent(e)
	case "UserCreatedEvent":
		e := events.UnsafeDeserializeAny[events.UserCreatedEvent](ev.Data)
		ua.apply_UserCreatedEvent(e)
	default:
		panic(fmt.Errorf("NO KNOWN EVENT %v", ev))
	}
	if ev.Version == 0 {
		ev.IsNew = true
	}
	v := len(ua.events) + 1 //for validation we need to start at 1
	ev.Version = int64(v)
	ev.AggregateType = ua.aggregateType
	ev.AggregateId = ua.aggregateId.String()
	ua.events = append(ua.events, *ev)
}

func (ua *UserAggregate) saveChanges() error {
	return ua.client.AddEvents(ua.aggregateId.String(), ua.events)
}
func (ua *UserAggregate) ChangeDisplayName(name string) error {
	if len(ua.events) == 0 {
		return fmt.Errorf("user does not yet exist")
	}

	if ua.displayName != name && len(name) <= 50 && time.Since(ua.changeDate).Seconds() > 10 {
		ua.addEvent(events.NewNameChangedEvent(name))
		err := ua.saveChanges()
		if err != nil {
			return fmt.Errorf("error saving changes")
		}
		return nil
	}
	return fmt.Errorf("validating username failed")
}

func (ua *UserAggregate) CreateUser(userCreate m.UserCreate) error {

	if len(ua.events) != 0 {
		return fmt.Errorf("user already exists")
	}

	if !strings.Contains(userCreate.Email, "@") {
		return fmt.Errorf("email does not contain @ sign")

	}
	if !(len(userCreate.Name) > 5 && len(userCreate.Name) < 50) {
		return fmt.Errorf("username not between 5 and 50 characters")
	}
	if !(len(userCreate.Password) >= 8 && len(userCreate.Password) < 50) {

		return fmt.Errorf("password not between 8 and 50 characters")

	}
	ua.addEvent(events.NewUserCreateEvent(userCreate))
	err := ua.saveChanges()
	if err != nil {
		return fmt.Errorf("ERROR ")
	}
	return nil
}
