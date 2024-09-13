package main

import (
	"math"
	"os"

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

	err = repository.AddEvent(&models.Event{Name: "myevent", Version: int64(math.Pow(2, 57)) - 1, Data: []byte{1, 2, 3}})

	if err != nil {
		log.Error().Err(err).Msg("Inserting into events table")
	}
}
