package main

import (
	"fmt"
	"os"

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
	evclient.AddEvents("myaggregate4444", []models.Event{{Version: 2, Name: "asdasd2", Data: []byte{1, 2, 3}}})
	evs, _ := evclient.GetEventsOrdered("myaggregate4444")
	fmt.Println(evs)

	h.Stop()
}
