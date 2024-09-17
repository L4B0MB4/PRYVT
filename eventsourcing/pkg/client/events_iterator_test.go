package client_test

import (
	"testing"

	"github.com/L4B0MB4/PRYVT/eventsourcing/pkg/client"
	"github.com/L4B0MB4/PRYVT/eventsourcing/pkg/models"
)

func TestIteratorNext(t *testing.T) {

	eIter := client.NewEventIterator([]models.Event{{}})
	ev, ok := eIter.Next()
	if ok != true || ev == nil {
		t.Error("SHOULD HAVE ONE EVENT")
		t.Fail()
	}
	ev, ok = eIter.Next()
	if ok == true || ev != nil {
		t.Error("SHOULD HAVE ONLY ONE EVENT")
		t.Fail()
	}
}

func TestIteratorNextMultiItem(t *testing.T) {

	eIter := client.NewEventIterator([]models.Event{{Name: "1"}, {Name: "2"}})
	ev, _ := eIter.Next()
	if ev.Name != "1" {
		t.Error("SHOULD HAVE FIRST EVENT IN FIRST NEXT CALL")
		t.Fail()
	}
	ev, _ = eIter.Next()
	if ev.Name != "2" {
		t.Error("SHOULD HAVE SECOND EVENT IN SECOND NEXT CALL")
		t.Fail()
	}
}

func TestIteratorCurrent(t *testing.T) {

	eIter := client.NewEventIterator([]models.Event{{Name: "1"}, {Name: "2"}})
	ev := eIter.Current()
	if ev != nil {
		t.Error("CURRENT SHOULD BE NULL BEFORE NEXT CALL")
	}
	eIter.Next()
	ev = eIter.Current()
	if ev.Name != "1" {
		t.Error("CURRENT SHOULD BE FIRST EVENT IN FIRST CALL")
		t.Fail()
	}
	ev = eIter.Current()
	if ev.Name != "1" {
		t.Error("CURRENT SHOULD STILL BE FIRST EVENT IN SECOND CALL")
		t.Fail()
	}
}
