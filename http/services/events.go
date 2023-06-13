package services

import "service-template-golang/domain/entity"

// EventsService implementation of interfaces.
type EventsService interface {
	GetID(ID string) (*entity.Events, error)
}
