package impl

import (
	"service-template-golang/domain/entity"
	"service-template-golang/http/repository"
)

// EventsServiceImpl encapsulates all the data necessary for the implementation of the EventsRepository.
type EventsServiceImpl struct {
	eventRepository repository.EventsRepository
}

// NewEventsService instance the repository services.
func NewEventsService(er repository.EventsRepository) *EventsServiceImpl {
	return &EventsServiceImpl{
		eventRepository: er,
	}
}

// GetID return the event by ID.
func (es *EventsServiceImpl) GetID(ID string) (*entity.Events, error) {
	return es.eventRepository.GetID(ID)
}
