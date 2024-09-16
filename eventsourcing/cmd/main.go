package main

import (
	"os"
	"time"

	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/client"
	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/httphandler"
	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/httphandler/controller"
	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/models"
	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	db := store.DatabaseConnection{}
	db.SetUp()
	conn, err := db.GetDbConnection()
	if err != nil {
		log.Error().Err(err).Msg("Unsuccessfull initalization of db")
		return
	}
	log.Debug().Msg("Db Connection was successful")
	repository := store.NewEventRepository(conn)

	c := controller.NewEventController(repository)
	h := httphandler.NewHttpHandler(c)

	go func() {
		h.Start()
	}()
	evclient, err := client.NewEventSourcingHttpClient("http://localhost:5515")
	if err != nil {
		panic(err)
	}
	evclient.AddEvents("myaggregate4444", []models.Event{{Version: 1, Name: "asdasd", Data: []byte{1, 2, 3}}})

	/*
		err = repository.AddEvent(&models.Event{AggregateType: "myaggregate", Name: "myevent", Version: 1, Data: []byte("erstes event")})

		if err != nil {
			log.Error().Err(err).Msg("Inserting into events table")
		}
		err = repository.AddEvent(&models.Event{AggregateType: "myaggregate", Name: "myevent2", Version: 2, Data: []byte("zweites event")})

		if err != nil {
			log.Error().Err(err).Msg("Inserting into events table")
		}
		err = repository.AddEvent(&models.Event{AggregateType: "myotheraggregate", Name: "whatanevent", Version: 1, Data: []byte("whatanevent")})

		if err != nil {
			log.Error().Err(err).Msg("Inserting into events table")
		}

		ev, err := repository.GetEventsForAggregate("myaggregate")
		if err != nil {
			log.Error().Err(err).Msg("myaggregate")
		}
		fmt.Println(ev)

		ev, err = repository.GetEventsForAggregate("myotheraggregate")
		if err != nil {
			log.Error().Err(err).Msg("myotheraggregate")
		}
		fmt.Println(ev)*/
	time.Sleep(100000)
}
