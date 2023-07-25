package events

import (
	"service-worker-sqs-postgres/core/domain"
	"service-worker-sqs-postgres/core/domain/entity"
	"service-worker-sqs-postgres/dataproviders/mapper"
	repository "service-worker-sqs-postgres/dataproviders/postgres/repository/events"
)

type IEventCaseUses interface {
	GetID(ID string) (*entity.Events, error)
}

// EventCaseUses encapsulates all the data necessary for the implementation of the EventsRepository.
type EventCaseUses struct {
	eventRepository repository.IEventRepository
}

// NewEventUseCases instance the repository usecases.
func NewEventUseCases(er repository.IEventRepository) *EventCaseUses {
	return &EventCaseUses{
		eventRepository: er,
	}
}

// GetID return the event by ID.
func (es *EventCaseUses) GetID(ID string) (*domain.Events, error) {
	event, err := es.eventRepository.GetID(ID)
	if err != nil {
		return nil, err
	}
	return mapper.ToDomainEvents(event), nil
}
