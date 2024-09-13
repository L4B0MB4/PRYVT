package models

type Event struct {
	Version       int64
	Name          string
	Data          []byte
	AggregateType string
}
