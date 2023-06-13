package domain

import "go.uber.org/zap"

// Event represents a process.
type Event struct {
	ID            string
	Retry         string
	Records       Events
	OriginalEvent interface{}
	Log           *zap.SugaredLogger
}

// Source represents a source of events.
type Source interface {
	Consume() <-chan *Event
	Processed(e *Event) error
	Close() error
}
