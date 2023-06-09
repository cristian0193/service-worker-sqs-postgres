package processor

import (
	"fmt"
	"go.uber.org/zap"
	"service-template-golang/domain"
	"time"
)

type Processor struct {
	Logger *zap.SugaredLogger
	Source domain.Source
}

func New(logger *zap.SugaredLogger, source domain.Source) (*Processor, error) {
	return &Processor{
		Logger: logger,
		Source: source,
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
	start := time.Now()
	key := fmt.Sprintf("%s", event.ID)

	for _, record := range event.Records {
		fmt.Print(key, " - ", record)
	}

	// notify that an event of consolidate is processed.
	p.Source.EventProcessed()
	event.Log.Infof("Consolidate event finished in %v", time.Since(start))
}

// Stop stops the Processor execution.
func (p *Processor) Stop() error {
	if err := p.Source.Close(); err != nil {
		return err
	}
	return nil
}
