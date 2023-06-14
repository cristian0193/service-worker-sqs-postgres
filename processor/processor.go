package processor

import (
	"go.uber.org/zap"
	"service-worker-sqs-postgres/database"
	"service-worker-sqs-postgres/domain"
	"time"
)

// Processor represents a process.
type Processor struct {
	logger   *zap.SugaredLogger
	source   domain.Source
	clientDB *database.ClientDB
}

// New instance a new processor.
func New(logger *zap.SugaredLogger, source domain.Source, db *database.ClientDB) (*Processor, error) {
	return &Processor{
		logger:   logger,
		source:   source,
		clientDB: db,
	}, nil
}

// Start a processor execution.
func (p *Processor) Start() {
	p.logger.Info("Starting processor")
	stream := p.source.Consume()
	for event := range stream {
		go p.handleEvent(event)
	}
}

// handleEvent is the entry point to handle consolidate event.
func (p *Processor) handleEvent(event *domain.Event) {
	if err := p.source.Processed(event); err != nil {
		event.Log.Errorf("Error processing event: %v", err)
	}
	elapsed := time.Since(time.Now())
	event.Log.Infof("Step 5 - Event finished in %dms", elapsed.Milliseconds())
}

// Stop stops the Processor execution.
func (p *Processor) Stop() error {
	return p.source.Close()
}
