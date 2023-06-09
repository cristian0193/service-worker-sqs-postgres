package domain

import "go.uber.org/zap"

// Event represents a new file to process.
type Event struct {
	ID            string
	Retry         string
	Records       []map[string]interface{}
	OriginalEvent interface{}
	Log           *zap.SugaredLogger
}

// Source represents a source of events.
type Source interface {
	Consume() <-chan *Event
	EventProcessed()
	Close() error
}
