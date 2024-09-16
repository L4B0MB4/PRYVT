package store_test

import (
	"os"
	"testing"

	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/models"
	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/store"
)

func setup() *store.DatabaseConnection {
	db := store.DatabaseConnection{}
	if _, err := os.Stat(store.GetDbFileLocation()); err == nil {
		db.Teardown()
	}
	db.SetUp()
	if !db.IsInitialized() {
		panic("ERROR DURING DB INIT ")
	}
	return &db

}
func teardown(db *store.DatabaseConnection) {
	err := db.Teardown()
	if err != nil {
		panic(err)
	}

}

func TestAddEventSuccessful(t *testing.T) {
	db := setup()
	defer teardown(db)
	conn, err := db.GetDbConnection()
	if err != nil {
		t.Error("CONNECTION SHOULD BE RETRIEVED WITHOUT A PROBLEM")
		t.Fail()
	}
	r := store.NewEventRepository(conn)
	ev := &models.Event{
		Version:     1,
		Name:        "testevent",
		Data:        []byte{0, 1},
		AggregateId: "anyaggregatetype",
	}
	err = r.AddEvent(ev)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	q, _ := conn.Query("SELECT * FROM events")
	if q.Next() != true {
		t.Error("THERE SHOULD BE ONE ENTRY IN THE DB")
		t.Fail()
	}
	if q.Next() == true {
		t.Error("THERE ARE MORE THEN ONE ENTRY IN THE DB")
		t.Fail()
	}

	evs, _ := r.GetEventsForAggregate(ev.AggregateId)
	evToComp := evs[0]
	if evToComp.AggregateId != ev.AggregateId || evToComp.Data[0] != ev.Data[0] || evToComp.Data[1] != ev.Data[1] || evToComp.Name != ev.Name || evToComp.Version != ev.Version {
		t.Error("SOMETHING WENT WRONG DURING SERIALIZATION OR DESERIALIZATION")
		t.Fail()
	}

}

func TestAddEventDuplicate(t *testing.T) {
	db := setup()
	defer teardown(db)
	conn, err := db.GetDbConnection()
	if err != nil {
		t.Error("CONNECTION SHOULD BE RETRIEVED WITHOUT A PROBLEM")
		t.Fail()
	}
	r := store.NewEventRepository(conn)
	ev := &models.Event{
		Version:     1,
		Name:        "testevent",
		Data:        []byte{0, 1},
		AggregateId: "anyaggregatetype",
	}
	err = r.AddEvent(ev)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	err = r.AddEvent(ev)
	if err == nil {
		t.Error("NO ERROR WHEN ADDING THE SAME EVENT TWICE")
		t.Fail()
	}
}

func TestAddTwoFollowingEvents(t *testing.T) {
	db := setup()
	defer teardown(db)
	conn, err := db.GetDbConnection()
	if err != nil {
		t.Error("CONNECTION SHOULD BE RETRIEVED WITHOUT A PROBLEM")
		t.Fail()
	}
	r := store.NewEventRepository(conn)
	ev := &models.Event{
		Version:     1,
		Name:        "testevent",
		Data:        []byte{0, 1},
		AggregateId: "anyaggregatetype",
	}
	err = r.AddEvent(ev)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	ev.Version++
	err = r.AddEvent(ev)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
func TestAddThreeEventsOfTwoAggregates(t *testing.T) {
	db := setup()
	defer teardown(db)
	conn, err := db.GetDbConnection()
	if err != nil {
		t.Error("CONNECTION SHOULD BE RETRIEVED WITHOUT A PROBLEM")
		t.Fail()
	}
	r := store.NewEventRepository(conn)
	oldAggType := "anyaggregatetype"
	ev := &models.Event{
		Version:     1,
		Name:        "testevent",
		Data:        []byte{0, 1},
		AggregateId: oldAggType,
	}
	r.AddEvent(ev)
	ev.Version++
	r.AddEvent(ev)
	ev.Version = 0
	newAggType := "aggregatetype2"
	ev.AggregateId = newAggType
	r.AddEvent(ev)

	i, _ := r.GetEventsForAggregate(oldAggType)
	if len(i) != 2 {
		t.Error("SHOULD HAVE 2 EVENTS FOR THIS AGGREGATE")
		t.Fail()
	}
	i, _ = r.GetEventsForAggregate(newAggType)
	if len(i) != 1 {
		t.Error("SHOULD HAVE 1 EVENT FOR THIS AGGREGATE")
		t.Fail()
	}
}
