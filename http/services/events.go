package services

import "service-worker-sqs-postgres/domain/entity"

// EventsService implementation of interfaces.
type EventsService interface {
	GetID(ID string) (*entity.Events, error)
}
