package events

import (
	"service-worker-sqs-postgres/core/domain/entity"
	"service-worker-sqs-postgres/dataproviders/repository/events"
)

type IEventsCaseUses interface {
	GetID(ID string) (*entity.Events, error)
}

// EventsCaseUses encapsulates all the data necessary for the implementation of the EventsRepository.
type EventsCaseUses struct {
	eventRepository events.IEventsRepository
}

// NewEventsUseCases instance the repository usecases.
func NewEventsUseCases(er events.IEventsRepository) *EventsCaseUses {
	return &EventsCaseUses{
		eventRepository: er,
	}
}

// GetID return the event by ID.
func (es *EventsCaseUses) GetID(ID string) (*entity.Events, error) {
	return es.eventRepository.GetID(ID)
}
