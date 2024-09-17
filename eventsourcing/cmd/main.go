package main

import (
	"os"

	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/httphandler"
	"gihtub.com/L4B0MB4/PRYVT/eventsouring/pkg/httphandler/controller"
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

	h.Start()
}
