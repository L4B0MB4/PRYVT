package models

type Event struct {
	Version     int64
	AggregateId int64
	Name        string
	Data        []byte
}
