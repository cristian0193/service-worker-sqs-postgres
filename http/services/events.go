package services

import (
	"service-template-golang/domain/entity"
	"service-template-golang/http/repository"
)

type EventsService struct {
	eventRepository *repository.EventsRepository
}

func NewEventsService(er *repository.EventsRepository) *EventsService {
	return &EventsService{
		eventRepository: er,
	}
}

func (es *EventsService) GetID(ID string) (*entity.Events, error) {
	return es.eventRepository.GetID(ID)
}
