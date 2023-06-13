package processor

import (
	"go.uber.org/zap"
	"service-template-golang/database"
	"service-template-golang/domain"
	"time"
)

type Processor struct {
	Logger   *zap.SugaredLogger
	Source   domain.Source
	ClientDB *database.ClientDB
}

func New(logger *zap.SugaredLogger, source domain.Source, rds *database.ClientDB) (*Processor, error) {
	return &Processor{
		Logger:   logger,
		Source:   source,
		ClientDB: rds,
	}, nil
}

// Start a processor execution.
func (p *Processor) Start() {
	p.Logger.Info("Starting processor")
	stream := p.Source.Consume()
	for event := range stream {
		go p.handleEvent(event)
	}
}

// handleEvent is the entry point to handle consolidate/upload event.
func (p *Processor) handleEvent(event *domain.Event) {
	event.Log.Infof("event started")
	if err := p.Source.Processed(event); err != nil {
		event.Log.Errorf("Event processed. %v", err)
	}
	event.Log.Infof("Consolidate event finished in %v", time.Since(time.Now()))
}

// Stop stops the Processor execution.
func (p *Processor) Stop() error {
	return p.Source.Close()
}
